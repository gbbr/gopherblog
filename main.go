package main

import (
	"github.com/backslashed/gopherblog/models"
	"github.com/backslashed/gopherblog/views"
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

	http.HandleFunc("/", views.ViewHome)
	http.HandleFunc("/post/", views.ViewPost)
	http.HandleFunc("/login", views.ViewLogin)
	http.HandleFunc("/edit", authenticate(views.ViewEdit))
	http.HandleFunc("/edit/", authenticate(views.ViewEditPost))

	err := http.ListenAndServe(conf.Host, nil)
	if err != nil {
		log.Fatal("Error starting server.")
	}
}
