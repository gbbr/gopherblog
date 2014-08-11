package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

const (
	SQL_POST_BY_ID = `
		SELECT slug, title, body, date, idUser, users.name, users.email, draft
		FROM posts 
		INNER JOIN users USING(idUser)
		WHERE idPost=?`
	SQL_POST_BY_SLUG = `
		SELECT slug, title, body, date, idUser, users.name, users.email, draft
		FROM posts 
		INNER JOIN users USING(idUser)
		WHERE slug=?`
	SQL_POSTS_BY_USER = `
		SELECT idPost, title, slug, date, draft
		FROM posts
		WHERE idUser=?
		ORDER BY draft DESC, date DESC`
	SQL_ALL_POSTS = `
		SELECT slug, title, date, idUser, users.name
		FROM posts
		INNER JOIN users USING(idUser)
		WHERE draft=false
		ORDER BY date DESC LIMIT ?`
	SQL_INSERT_POST = `
		INSERT INTO posts (slug, title, body, idUser, draft)
		VALUES (?, ?, ?, ?, ?)`
	SQL_UPDATE_POST = `
		UPDATE posts SET slug=?, title=?, body=?, idUser=?, draft=?
		WHERE idPost=?`

	SQL_USER_BY_ID = `
		SELECT name, email 
		FROM users 
		WHERE idUser=?`
	SQL_USER_AUTH = `
		SELECT name, idUser
		FROM users 
		WHERE email=? AND password=?`
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
