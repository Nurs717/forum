package controller

import (
	"net/http"
)

func successfulreg() http.HandlerFunc { //Y
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.ServeFile(w, r, "view/successfulreg.html")
		}
	}
}
