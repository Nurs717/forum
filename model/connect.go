package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

// Connect connecting func
func Connect() *sql.DB {

	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db
	return db

}
