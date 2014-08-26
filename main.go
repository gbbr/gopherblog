package main

import (
	"flag"
	"github.com/gbbr/gopherblog/controllers"
	"github.com/gbbr/gopherblog/models"
	"log"
	"net/http"
)

var (
	host     = flag.String("host", "localhost", "Hostname for HTTP server")
	port     = flag.String("port", "8080", "Port for HTTP server")
	dbString = flag.String("db", "root@tcp(localhost:3306)/gopherblog", "Database connection string")
)

func main() {
	flag.Parse()

	models.ConnectDb(*dbString)
	defer models.CloseDb()

	mux := http.NewServeMux()

	// Blog pages
	mux.HandleFunc("/", controller.Posts)
	mux.HandleFunc("/post/", controller.Post)
	mux.HandleFunc("/tag/", controller.Posts)

	// Admin routes
	mux.HandleFunc("/login", controller.Login)
	mux.HandleFunc("/new", authenticate(controller.NewPost))
	mux.HandleFunc("/manage", authenticate(controller.Manage))
	mux.HandleFunc("/edit/", authenticate(controller.EditPost))

	// Static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("/bower_components/", http.StripPrefix("/bower_components/", http.FileServer(http.Dir("bower_components"))))

	log.Printf("Starting on %s:%s\n", *host, *port)
	err := http.ListenAndServe(*host+":"+*port, mux)

	log.Fatal(err)
}
