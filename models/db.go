package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

const (
	SQL_POST_BY_ID = `
		SELECT title, body, date, idUser, users.name, users.email 
		FROM posts 
		INNER JOIN users USING(idUser)
		WHERE idPost=?`
	SQL_POST_BY_SLUG = `
		SELECT title, body, date, idUser, users.name, users.email 
		FROM posts 
		INNER JOIN users USING(idUser)
		WHERE slug=?`
	SQL_ALL_POSTS = `
		SELECT slug, title, date, idUser, users.name
		FROM posts
		INNER JOIN users USING(idUser)
		ORDER BY date DESC LIMIT ?`

	SQL_USER_BY_ID = `
		SELECT name, email FROM users WHERE idUser=?`
	SQL_USER_AUTH = `
		SELECT name, idUser, isAuthor FROM users WHERE email=? AND password=?`
)

// Creates and tests database connection
func ConnectDb(address string) {
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
