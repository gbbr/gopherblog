package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db *sql.DB

type Config struct {
	Host     string
	DbString string
}

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
