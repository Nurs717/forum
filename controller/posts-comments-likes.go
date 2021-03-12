package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"forum/model"
	"forum/view"
)

//Flag flag
var Flag bool
var filter string
var myposts string
var posts []view.Post

//UserID userid
var UserID int

//UserName username
var UserName string

func postsandlikes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates := template.Must(template.ParseGlob("view/*.html"))

		if myposts != "" {
			posts = (model.GetPosts(myposts, UserID))
		} else if filter != "" {
			posts = (model.GetPosts(filter, UserID))
		} else {
			posts = model.GetPosts("", UserID)
		}

		if r.Method == "GET" {
			fmt.Println("MethodGet")
			cookie, err := r.Cookie("session")

			resetFilter := r.FormValue("resetfilter")
			if resetFilter == "resetfilter" {
				posts = model.GetPosts("", UserID)
			}

			logout := r.FormValue("logout")
			fmt.Println("This is logout:", logout)
			if logout == "logout" {
				cookie.Expires = time.Now().AddDate(0, 0, -1)
				http.SetCookie(w, cookie)
				fmt.Println("Loged Out")
				http.Redirect(w, r, "/", http.StatusFound)
			}

			login := r.FormValue("login")
			fmt.Println("Formvalue Login:", login)
			if login == "login" {
				fmt.Println("Redirected to login page")
				http.Redirect(w, r, "/", http.StatusFound)
			}

			if err != nil {
				Flag = false
			} else {
				Flag = model.IsUserValid(cookie.Value)
				UserID, UserName = model.GetUserIDbySession(cookie.Value)
			}
			response := struct {
				Username string
				Posts    []view.Post
			}{
				Username: UserName,
				Posts:    posts,
			}

			if Flag {
				if err := templates.ExecuteTemplate(w, "posts.html", response); err != nil {
					http.Error(w, "Internal Server Error!!!\nERROR-500", http.StatusInternalServerError)
					fmt.Println(filter)
					return
				}

			} else {
				if err := templates.ExecuteTemplate(w, "postsonly.html", posts); err != nil {
					http.Error(w, "Internal Server Error!!!\nERROR-500", http.StatusInternalServerError)
					return
				}

			}
			fmt.Println("EndMethodGet")
		}

		if r.Method == "POST" {
			fmt.Println("Start MethodPost")
			cookie, err := r.Cookie("session")
			if err != nil && err != http.ErrNoCookie {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			fmt.Println("ya proshla dalshe")

			if cookie != nil {
				UserID, UserName = model.GetUserIDbySession(cookie.Value)
				fmt.Println("UserID and UserName recieved", UserID, UserName)
				Flag = model.IsUserValid(cookie.Value)
				if !Flag {
					http.Error(w, "StatusUnauthorized!!\nERROR-401", http.StatusUnauthorized)
					templates.ExecuteTemplate(w, "postsonly.html", posts)
					return
				}
			}

			// logo := []view.User{}

			myposts = r.FormValue("myposts")
			fmt.Println("This mypost:", myposts)
			if myposts != "" {
				fmt.Println(model.GetPosts(myposts, UserID))
			}
			filter = r.FormValue("filter")
			fmt.Println("This filter:", filter)
			post := r.FormValue("post")
			category := r.FormValue("category")

			fmt.Println("This post:", post)
			fmt.Println("This category:", category)

			if len(post) > 0 {
				if err := model.AddPost(UserID, post, UserName, category); err != nil {
					log.Fatal(err)
				}
			}
			post = ""

			PostIDL, _ := strconv.Atoi(r.FormValue("PostID"))
			fmt.Println("This Post ID:", PostIDL)

			comment := r.FormValue("comment")
			PostIDC, _ := strconv.Atoi(r.FormValue("PostIDC"))
			CommentID, _ := strconv.Atoi(r.FormValue("CommentID"))

			fmt.Println("This post IDC", PostIDC)

			if PostIDL != 0 {

				if model.IsLiked(UserID, PostIDL) == false {
					if err := model.AddLike(UserID, PostIDL); err != nil {
						log.Fatal(err)
					}
				} else {

					fmt.Println("Like is deleted")

				}
			}
			if CommentID != 0 {

				if model.IsCommentLiked(CommentID, UserID) == false {
					if err := model.AddCommentLike(CommentID, UserID); err != nil {
						log.Fatal(err)
					}
				} else {

					fmt.Println("Comment Like is deleted")

				}
			}
			CommentID = 0
			if PostIDC != 0 {
				if err := model.AddComment(UserID, PostIDC, comment, UserName); err != nil {
					log.Fatal(err)
				}
				fmt.Println(comment)
				fmt.Println(PostIDC)
			}
			fmt.Println("End PostAndLikes")

			http.Redirect(w, r, "/posts", 302)

		}
	}
}
