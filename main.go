package main

import (
	"github.com/gbbr/gopherblog/controllers"
	"github.com/gbbr/gopherblog/models"
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
	defer models.CloseDb()

	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/post/", controller.Post)
	mux.HandleFunc("/login", controller.Login)
	mux.HandleFunc("/new", authenticate(controller.NewPost))
	mux.HandleFunc("/manage", authenticate(controller.Manage))
	mux.HandleFunc("/edit/", authenticate(controller.EditPost))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(conf.Host, mux)

	log.Fatal(err)
}
