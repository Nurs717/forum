package model

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Insert adding new user
func Insert(a, b, c string) error {

	hashed, _ := bcrypt.GenerateFromPassword([]byte(c), bcrypt.DefaultCost)
	password := string(hashed)
	_, err := con.Exec("INSERT INTO User (Email,User_Name,Password) VALUES(?,?,?)", b, a, password)
	if err != nil {
		return err
	}
	fmt.Println("Data is inserted, insert")
	return nil

}

// AddSession adding session
func AddSession(a, b string) error {

	_, err := con.Exec("INSERT INTO Session (Email,Session_Cookie) VALUES(?,?)", a, b)
	if err != nil {
		fmt.Println("Err, add session exist")
		return err
	}
	fmt.Println("Data is inserted, add session")
	return nil

}

// AddPost handling all posts page
func AddPost(a int, b string, UserName string, category string) error {

	t := time.Now()
	fmt.Println(t.String())
	T := t.Format("2006-01-02 15:04:05")
	date := T[0:11]
	time := T[11:]
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	_, err := con.Exec("INSERT INTO Post (User_ID, Post_Body, Post_Date, Post_Time, User_Name, Categories) VALUES(?,?,?,?,?,?)", a, b, date, time, UserName, category)
	if err != nil {
		fmt.Println("Error AddPost")
		return err
	}
	fmt.Println("Data is inserted, add post")
	return nil

}

// AddLike adding like
func AddLike(a int, b int) error {

	_, err := con.Exec("INSERT INTO Like (User_ID, Post_ID) VALUES(?,?)", a, b)
	if err != nil {
		fmt.Println("Error insert like post")
		return err
	}
	fmt.Println("Data is inserted, add like")
	return nil

}

// AddCommentLike adding comment like
func AddCommentLike(a int, b int) error {

	_, err := con.Exec("INSERT INTO Like_Comment (Comment_ID, User_ID) VALUES(?,?)", a, b)
	if err != nil {
		fmt.Println("Error Add comment like")
		return err
	}
	fmt.Println("Data is inserted, comment like")
	return nil

}

// AddComment adding comment
func AddComment(a int, b int, comment string, UserName string) error {
	_, err := con.Exec("INSERT INTO Comment (User_ID, Post_ID, Comment, User_Name) VALUES(?,?,?,?)", a, b, comment, UserName)
	if err != nil {
		fmt.Println("Error add comment")
		return err
	}
	fmt.Println("Data is inserted, add comment")
	return nil

}
