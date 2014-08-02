package main

import (
	"github.com/backslashed/gopherblog/controllers"
	"github.com/backslashed/gopherblog/models"
	"log"
	"net/http"
)

// Holds web server and database connection strings
type Config struct {
	Host     string
	DbString string
}

func main() {
	conf := Config{
		Host:     "mecca.local:8080",
		DbString: "root:root@tcp(localhost:3306)/blog",
	}

	models.ConnectDb(conf.DbString)

	http.HandleFunc("/", controller.ViewHome)
	http.HandleFunc("/post/", controller.ViewPost)
	http.HandleFunc("/login", controller.ViewLogin)
	http.HandleFunc("/edit", authenticate(controller.ViewEdit))
	http.HandleFunc("/edit/", authenticate(controller.ViewEditPost))

	err := http.ListenAndServe(conf.Host, nil)
	if err != nil {
		log.Fatal("Error starting server.")
	}
}
