package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Implicit import. database/sql package will use it.
)

func main() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))

	db, error := sql.Open("mysql", connectionString)

	if error != nil {
		fmt.Println("Error at sql.Open")
		log.Fatal(error)
	}
	defer db.Close() // After all, close the connection

	// check if connection is open
	if error = db.Ping(); error != nil {
		fmt.Println("Error on ping")
		log.Fatal(error)
	}

	fmt.Println("Connection opened.")

	lines, error := db.Query("SELECT * FROM USERS")
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(lines)
}
