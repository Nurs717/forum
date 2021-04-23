package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"

	"forum/model"
)

func login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {

			if r.URL.Path != "/" {
				http.Error(w, "Error 404", http.StatusNotFound)
			}

			if r.URL.Path == "/" {
				http.ServeFile(w, r, "view/main.html")
			}
		}

		if r.Method == "POST" {
			email := r.FormValue("email")

			password := r.FormValue("password")

			if model.IsValid(email, password) == true {

				cookie, err := r.Cookie("session")
				if err == nil {
					http.Redirect(w, r, "/posts", 302)
					fmt.Println("exist")
				}

				if err != nil {

					expire := time.Now().Add(24 * time.Hour)

					u1 := uuid.NewV4()

					cookie = &http.Cookie{
						Name:     "session",
						Value:    u1.String(),
						Expires:  expire,
						Path:     "/",
						HttpOnly: true,
					}

					http.SetCookie(w, cookie)

					if err := model.AddSession(email, cookie.Value); err != nil {
						if err := model.UpdateSession(cookie.Value, email); err != nil {
							log.Fatal(err)
						}
					}
					fmt.Println("new user set")
					http.Redirect(w, r, "/posts", 302)
				}

				return

			}

			http.Error(w, "Invalid email or password", http.StatusNotFound)

		}

	}
}
