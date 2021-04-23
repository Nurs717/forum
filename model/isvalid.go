package model

import (
	"fmt"
	"log"

	"forum/view"

	"golang.org/x/crypto/bcrypt"
)

// IsValid is user valid
func IsValid(a, b string) bool {

	rows, err := con.Query("select * from User")
	if err != nil {
		log.Fatal(err)
	}
	Users := []view.User{}
	for rows.Next() {
		u := view.User{}
		err := rows.Scan(&u.ID, &u.Email, &u.Username, &u.Password)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		Users = append(Users, u)
	}
	for _, u := range Users {
		if u.Email == a {
			password := []byte(b)
			hashed := []byte(u.Password)
			if err := bcrypt.CompareHashAndPassword(hashed, password); err == nil {
				return true
			}
		}
	}
	return false
}

// IsUserValid is session user valid
func IsUserValid(Session string) bool {

	rows, err := con.Query("select * from Session")
	if err != nil {
		log.Fatal(err)
	}
	Users := []view.SessionID{}
	for rows.Next() {
		s := view.SessionID{}
		err := rows.Scan(&s.Email, &s.SessionID)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		Users = append(Users, s)
	}
	for _, s := range Users {
		if s.SessionID == Session {
			fmt.Println("Checked uservalid")
			return true

		}
	}
	return false

}

// GetPosts posts
func GetPosts(filter string, UserID int) []view.Post {
	var post string

	rows, err := con.Query("select t1.*, count(t2.User_ID) as LikeCount from Comment t1 left join Like_Comment t2 USING(Comment_ID) group by t1.Comment_ID")
	if err != nil {
		log.Fatal(err)
	}
	Comments := []view.Comment{}
	for rows.Next() {
		p := view.Comment{}
		err := rows.Scan(&p.CommentID, &p.CommentBody, &p.UserID, &p.PostID, &p.UserName, &p.LikeCounts)

		if err != nil {
			fmt.Println("Error")
			continue
		}
		Comments = append(Comments, p)
	}
	fmt.Println("1")

	if filter == "sport" {
		post = "select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) where t1.Categories='sport' group by t1.Post_ID order by Post_ID DESC"
	} else if filter == "religion" {
		post = "select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) where t1.Categories='religion' group by t1.Post_ID order by Post_ID DESC"
	} else if filter == "politics" {
		post = "select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) where t1.Categories='politics' group by t1.Post_ID order by Post_ID DESC"
	} else if filter == "science" {
		post = "select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) where t1.Categories='science' group by t1.Post_ID order by Post_ID DESC"
	} else if filter == "others" {
		post = "select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) where t1.Categories='others' group by t1.Post_ID order by Post_ID DESC"
	} else {
		post = "select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) group by t1.Post_ID order by Post_ID DESC"
	}
	if filter == "myposts" {
		fmt.Println("myposts filter", UserID)
		rows, err := con.Query("select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) where t1.User_ID=? group by t1.Post_ID order by Post_ID DESC", UserID)
		if err != nil {
			log.Fatal(err)
		}
		Posters := []view.Post{}
		for rows.Next() {
			p := view.Post{}
			err := rows.Scan(&p.PostID, &p.UserID, &p.PostBody, &p.PostDate, &p.PostTime, &p.UserName, &p.Category, &p.LikeCounts)

			if err != nil {
				fmt.Println("Error my posts filter")
				continue
			}
			Posters = append(Posters, p)
		}
		fmt.Println("2")
		i := 0
		for _, p := range Posters {
			for _, c := range Comments {
				if p.PostID == c.PostID {
					p.Comments = append(p.Comments, c)
				}

			}
			Posters[i].Comments = p.Comments
			i++

		}

		return Posters
	}
	if filter == "myfavourite" {
		fmt.Println("myfavourite filter", UserID)
		rows, err := con.Query("select t1.*, count(t2.User_ID) from Post t1 left join Like t2 USING(Post_ID) where t2.User_ID=? group by t1.Post_ID order by Post_ID DESC", UserID)
		if err != nil {
			log.Fatal(err)
		}
		Posters := []view.Post{}
		for rows.Next() {
			p := view.Post{}
			err := rows.Scan(&p.PostID, &p.UserID, &p.PostBody, &p.PostDate, &p.PostTime, &p.UserName, &p.Category, &p.LikeCounts)

			if err != nil {
				fmt.Println("Error my favourite posts")
				continue
			}
			Posters = append(Posters, p)
		}
		fmt.Println("3")
		i := 0
		for _, p := range Posters {
			for _, c := range Comments {
				if p.PostID == c.PostID {
					p.Comments = append(p.Comments, c)
				}

			}
			Posters[i].Comments = p.Comments
			i++

		}

		return Posters
	}
	rows, err = con.Query(post)
	if err != nil {
		log.Fatal(err)
	}
	Posters := []view.Post{}
	for rows.Next() {
		p := view.Post{}
		err := rows.Scan(&p.PostID, &p.UserID, &p.PostBody, &p.PostDate, &p.PostTime, &p.UserName, &p.Category, &p.LikeCounts)

		if err != nil {
			fmt.Println("Error 4")
			continue
		}
		Posters = append(Posters, p)
	}
	fmt.Println("4")
	i := 0
	for _, p := range Posters {
		for _, c := range Comments {
			if p.PostID == c.PostID {
				p.Comments = append(p.Comments, c)
			}

		}
		Posters[i].Comments = p.Comments
		i++

	}
	fmt.Println("EndGetPost")
	return Posters

}

// GetUserIDbySession gets ID by session
func GetUserIDbySession(ID string) (int, string) {
	UserName := "..."
	UserID := 0
	rows, err := con.Query("select * from Session")
	if err != nil {
		log.Fatal(err)
	}
	Users := []view.SessionID{}
	for rows.Next() {
		s := view.SessionID{}
		err := rows.Scan(&s.Email, &s.SessionID)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		Users = append(Users, s)

	}
	for _, s := range Users {
		if s.SessionID == ID {
			UserID, UserName = GetUsrIDandName(s.Email)
			fmt.Println("GetUserIDbySession")

		}
	}
	return UserID, UserName

}

// GetUsrIDandName gets user ID and name
func GetUsrIDandName(Email string) (int, string) {

	UserName := ""
	UserID := 0
	rows, err := con.Query("select * from User")
	if err != nil {
		log.Fatal(err)
	}
	Users := []view.User{}
	for rows.Next() {
		u := view.User{}
		err := rows.Scan(&u.ID, &u.Email, &u.Username, &u.Password)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		Users = append(Users, u)
	}
	for _, u := range Users {
		if u.Email == Email {
			UserID = u.ID
			UserName = u.Username
			fmt.Println("GetUserIdandName")
		}
	}
	return UserID, UserName

}

// IsLiked checks is post liked
func IsLiked(UserID int, PostID int) bool {

	rows, err := con.Query("select * from Like")
	if err != nil {
		log.Fatal(err)
	}

	Users := []view.Like{}

	for rows.Next() {
		s := view.Like{}
		err := rows.Scan(&s.UserID, &s.PostID)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		Users = append(Users, s)

	}

	for _, s := range Users {
		if s.UserID == UserID && s.PostID == PostID {
			_, err := con.Exec("delete from Like where User_ID=? and Post_ID=?", UserID, PostID)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Like is deleted")
			return true
		}
	}
	return false
}

// IsCommentLiked checks is comment liked
func IsCommentLiked(CommentID int, UserID int) bool {

	rows, err := con.Query("select * from Like_Comment")
	if err != nil {
		log.Fatal(err)
	}

	Users := []view.CommentLike{}

	for rows.Next() {
		s := view.CommentLike{}
		err := rows.Scan(&s.CommentID, &s.UserID)
		if err != nil {
			fmt.Println("Error")
			continue
		}
		Users = append(Users, s)

	}

	for _, s := range Users {
		if s.UserID == UserID && s.CommentID == CommentID {
			_, err := con.Exec("delete from Like_Comment where User_ID=? and Comment_ID=?", UserID, CommentID)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("CommentLike is deleted")
			return true
		}
	}
	return false
}
