package controller

import (
	"net/http"

	"github.com/mazhaboy/forum/tree/master/model"
)

func registration() http.HandlerFunc { //Y
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "view/registration.html")
		}
		if r.Method == "POST" {

			username := r.FormValue("username")
			email := r.FormValue("email")
			password := r.FormValue("password")

			if len(username) < 5 || len(email) < 5 || len(password) < 5 {
				http.Error(w, "Your email, username or password has to contain at least 5 characters!", http.StatusBadRequest)
				return
			}

			if err := model.Insert(username, email, password); err != nil {
				http.Error(w, "This email or username is already exists try another one!", http.StatusBadRequest)
			} else {
				http.Redirect(w, r, "/successfulreg", 302)
				return
			}

			http.ServeFile(w, r, "view/registration.html")

		}
	}
}
