package main

import (
	"flag"
	"github.com/gbbr/gopherblog/controllers"
	"github.com/gbbr/gopherblog/models"
	"log"
	"net/http"
)

// Command line flags
var (
	host     = flag.String("host", "localhost", "Hostname for HTTP server")
	port     = flag.String("port", "8080", "Port for HTTP server")
	dbString = flag.String("db", "root:root@tcp(localhost:3306)/gopherblog", "Database connection string, defaults to 'root:root@tcp(localhost:3306)/gopherblog'")
)

func main() {
	flag.Parse()

	models.ConnectDb(*dbString)
	defer models.CloseDb()

	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/post/", controller.Post)
	mux.HandleFunc("/login", controller.Login)
	mux.HandleFunc("/new", authenticate(controller.NewPost))
	mux.HandleFunc("/manage", authenticate(controller.Manage))
	mux.HandleFunc("/edit/", authenticate(controller.EditPost))

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(*host+":"+*port, mux)

	log.Fatal(err)
}
