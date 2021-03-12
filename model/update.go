package model

import (
	"fmt"
)

// UpdateSession update session
func UpdateSession(NewSessionID, Email string) error {

	_, err := con.Exec("update Session set SessionID=? where Email=?", NewSessionID, Email)
	if err != nil {
		return err
	}
	fmt.Println(Email)
	fmt.Println(NewSessionID)
	fmt.Println("ExpiredSessionID is updated")
	return nil

}

// UpdateLike update Like
func UpdateLike(Likes int, PostID string) error {

	_, err := con.Exec("update pl set Like=? where PostID=?", Likes, PostID)
	if err != nil {
		return err
	}
	fmt.Println(Likes)
	fmt.Println(PostID)
	fmt.Println("Likes is updated")
	return nil

}
