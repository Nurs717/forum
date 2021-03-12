package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mazhaboy/forum/tree/master/controller"
	"github.com/mazhaboy/forum/tree/master/model"
)

func main() {
	mux := controller.Register()
	fmt.Println("Starting server at :8282")
	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":8282", mux))
}
