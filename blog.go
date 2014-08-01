package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db *sql.DB

type Config struct {
	Host     string
	DbString string
}

func ViewPost(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/posts/")-1:]
	post := &Post{slug: slug}
	err := post.Fetch()
	if err != nil {
		fmt.Fprintf(w, "%+v", "Post not found")
	}
	fmt.Fprintf(w, "%+v", post)
}

func main() {
	var err error

	conf := Config{
		Host:     "localhost:8080",
		DbString: "root:root@tcp(localhost:3306)/blog",
	}

	http.HandleFunc("/post/", ViewPost)

	db, err = sql.Open("mysql", conf.DbString)
	if err != nil {
		log.Fatal("Error opening DB")
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error connecting to DB")
	}

	err = http.ListenAndServe(conf.Host, nil)
	if err != nil {
		log.Fatal("Error starting server.")
	}
}
