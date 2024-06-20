package services

import (
	"crud/dao"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func writeError(err error, errCode int, errDescription string, w *http.ResponseWriter) {
	wr := *w
	wr.WriteHeader(errCode)
	wr.Write([]byte(errDescription))
	fmt.Printf(errDescription+" - %s\n", err)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(err, http.StatusBadRequest, "Error reading request body", &w)
		return
	}

	var user User
	if err = json.Unmarshal(body, &user); err != nil {
		writeError(err, http.StatusBadRequest, "Malformed JSON", &w)
		return
	}

	db, err := dao.Getconnection()
	if err != nil {
		writeError(err, http.StatusServiceUnavailable, "Fail to connect to database", &w)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO USERS (NAME,AGE) VALUES (?,?)")
	if err != nil {
		writeError(err, http.StatusServiceUnavailable, "Fail to create db statement", &w)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Age)
	if err != nil {
		writeError(err, http.StatusServiceUnavailable, "Fail to execute statement", &w)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		writeError(err, http.StatusServiceUnavailable, "Fail to get id of new record", &w)
		return
	}

	w.Write([]byte(fmt.Sprintf("Created %d\n", id)))
	w.WriteHeader(http.StatusOK)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	dao, error := dao.Getconnection()
	if error != nil {
		writeError(error, 500, "Error getting db connection", &w)
		return
	}

	rows, error := dao.Query("SELECT * FROM USERS")
	if error != nil {
		writeError(error, 500, "Query failed", &w)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if error := rows.Scan(&user.Name, &user.Age, &user.Id); error != nil {
			writeError(error, 500, "Failed to scan results from database", &w)
			return
		}
		users = append(users, user)
	}

	if error := json.NewEncoder(w).Encode(users); error != nil {
		writeError(error, 500, "Failed to encode json", &w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, error := strconv.ParseUint(params["id"], 10, 32)
	if error != nil {
		writeError(error, http.StatusInternalServerError, "Failed to parse id", &w)
		return
	}

	db, error := dao.Getconnection()
	if error != nil {
		writeError(error, http.StatusInternalServerError, "Failed to get db connection", &w)
		return
	}

	defer db.Close()

	var user User
	row := db.QueryRow("select * from USERS where ID = ?", id)
	if error := row.Scan(&user.Name, &user.Age, &user.Id); error != nil {
		if error == sql.ErrNoRows {
			writeError(error, http.StatusNotFound, "No rows", &w)
		} else {
			writeError(error, http.StatusInternalServerError, "Failed on query", &w)
		}
		return
	}

	if error := json.NewEncoder(w).Encode(user); error != nil {
		writeError(error, http.StatusInternalServerError, "Failed on json encode", &w)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, error := strconv.ParseUint(params["id"], 10, 32)
	if error != nil {
		writeError(error, http.StatusInternalServerError, "Failed to parse id", &w)
		return
	}

	var user User
	if error := json.NewDecoder(r.Body).Decode(&user); error != nil {
		writeError(error, http.StatusInternalServerError, "Failed to decode body", &w)
		return
	}

	db, error := dao.Getconnection()
	if error != nil {
		writeError(error, http.StatusInternalServerError, "Failed to connect to db", &w)
		return
	}

	defer db.Close()

	stmt, error := db.Prepare("update USERS set NAME = ?, AGE = ? where ID = ?")
	if error != nil {
		writeError(error, http.StatusInternalServerError, "Failed to prepare query statement", &w)
		return
	}
	defer stmt.Close()

	result, error := stmt.Exec(user.Name, user.Age, id)
	if error != nil {
		writeError(error, http.StatusInternalServerError, "Failed to on updating", &w)
		return
	}

	affected, error := result.RowsAffected()
	if error != nil {
		writeError(error, http.StatusInternalServerError, "Failed result", &w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	fmt.Printf("Updated %d rows\n", affected)

}
