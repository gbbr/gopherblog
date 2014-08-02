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
		Host:     "localhost:8080",
		DbString: "root:root@tcp(localhost:3306)/blog",
	}

	models.ConnectDb(conf.DbString)
	http.HandleFunc("/post/", views.ViewPost)

	err := http.ListenAndServe(conf.Host, nil)
	if err != nil {
		log.Fatal("Error starting server.")
	}
}
