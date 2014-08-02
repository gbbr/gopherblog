package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db *sql.DB

// Holds web server and database connection strings
type Config struct {
	Host     string
	DbString string
}

// Creates and tests database connection
func connectDB(address string) {
	var err error
	db, err = sql.Open("mysql", address)
	if err != nil {
		log.Fatal("Error opening DB")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to DB")
	}
}

func main() {
	conf := Config{
		Host:     "localhost:8080",
		DbString: "root:root@tcp(localhost:3306)/blog",
	}

	http.HandleFunc("/post/", viewPost)

	connectDB(conf.DbString)

	err := http.ListenAndServe(conf.Host, nil)
	if err != nil {
		log.Fatal("Error starting server.")
	}
}
