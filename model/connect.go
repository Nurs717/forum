package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

// Connect connecting to Database
func Connect() *sql.DB {
	// var connection *sqlite3.SQLiteConn

	// sql.Register("sqlite3conn", &sqlite3.SQLiteDriver{ConnectHook: func(conn *sqlite3.SQLiteConn) error {
	// 	connection = conn
	// 	return nil
	// }})

	db, err := sql.Open("sqlite3", "file:forum.db?_auth&_auth_user=Nurs&_auth_pass=alemschool&_auth_crypt=sha1")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	tables := []string{`CREATE TABLE IF NOT EXISTS Comment (Comment_ID INTEGER PRIMARY KEY, Comment TEXT, User_ID INTEGER, Post_ID INTEGER, User_Name TEXT);`,
		`CREATE TABLE IF NOT EXISTS Like (User_ID INTEGER, Post_ID INTEGER);`,
		`CREATE TABLE IF NOT EXISTS Like_Comment (Comment_ID INTEGER, User_ID INTEGER);`,
		`CREATE TABLE IF NOT EXISTS Post (Post_ID INTEGER PRIMARY KEY, User_ID TEXT, Post_Body TEXT, Post_Date TEXT, Post_Time TEXT, User_Name TEXT, Categories TEXT);`,
		`CREATE TABLE IF NOT EXISTS Session (Email TEXT, Session_Cookie TEXT);`,
		`CREATE TABLE IF NOT EXISTS User (User_ID INTEGER PRIMARY KEY, Email TEXT, User_Name TEXT, Password TEXT);`}

	for _, v := range tables {
		_, err = db.Exec(v)
		if err != nil {
			log.Fatalf("ERROR EXEC: %q\n%v", v, err.Error())
		}
	}

	fmt.Println("Connected to the database")
	con = db
	return db
}
