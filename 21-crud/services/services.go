package services

import (
	"crud/dao"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
