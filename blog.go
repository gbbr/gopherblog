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

func viewPost(w http.ResponseWriter, r *http.Request) {
	post := &Post{
		slug: r.URL.Path[len("/posts/")-1:],
	}

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

	http.HandleFunc("/post/", viewPost)

	db, err = sql.Open("mysql", conf.DbString)
	if err != nil {
		log.Fatal("Error opening DB")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to DB")
	}

	err = http.ListenAndServe(conf.Host, nil)
	if err != nil {
		log.Fatal("Error starting server.")
	}
}
