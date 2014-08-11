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

	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/post/", controller.Post)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/new", authenticate(controller.NewPost))
	http.HandleFunc("/manage", authenticate(controller.Manage))
	http.HandleFunc("/edit/", authenticate(controller.EditPost))

	log.Fatal(http.ListenAndServe(conf.Host, nil))
	models.CloseDb()
}
