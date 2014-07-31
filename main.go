package main

import (
	"github.com/backslashed/gopherblog/helpers"
	"net/http"
)

const (
	HOSTNAME = "localhost"
	PORT     = "8080"
)

func main() {
	err := http.ListenAndServe(HOSTNAME+":"+PORT, nil)
	helpers.HandleError(err, "Error starting server")
}
