package controller

import "net/http"

// Register reg
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/registration", registration())
	mux.HandleFunc("/successfulreg", successfulreg())
	mux.HandleFunc("/posts", postsandlikes())
	mux.HandleFunc("/", login())
	return mux
}
