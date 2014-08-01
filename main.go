package main

import (
	"github.com/backslashed/gopherblog/blog"
	"net/http"
)

var config = blog.Config{
	Host:     "localhost",
	Port:     "8080",
	DbString: "root:root@tcp(localhost:3306)/blog",
}

func main() {
	err := blog.Start(config)
	blog.HandleError(err, "Error starting blog instance")

	http.Handle("/post/", new(blog.Post))

	err = http.ListenAndServe(config.Host+":"+config.Port, nil)
	blog.HandleError(err, "Error starting HTTP server")
}
