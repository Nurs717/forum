package model

import (
	"fmt"
)

// UpdateSession update session
func UpdateSession(NewSessionID, Email string) error {

	_, err := con.Exec("update Session set Session_Cookie=? where Email=?", NewSessionID, Email)
	if err != nil {
		return err
	}
	fmt.Println(Email)
	fmt.Println(NewSessionID)
	fmt.Println("SessionID is updated")
	return nil

}

// UpdateLike update Like
func UpdateLike(Likes int, PostID string) error {

	_, err := con.Exec("update pl set Like=? where Post_ID=?", Likes, PostID)
	if err != nil {
		return err
	}
	fmt.Println(Likes)
	fmt.Println(PostID)
	fmt.Println("Likes is updated")
	return nil

}
