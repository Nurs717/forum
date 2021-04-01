package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/controller"
	"forum/model"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	mux := controller.Register()
	fmt.Println("Starting server at :8282")
	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":8282", mux))
}
