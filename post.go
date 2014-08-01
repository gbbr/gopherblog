package main

import (
	"database/sql"
	"errors"
	mysql "github.com/go-sql-driver/mysql"
)

type Post struct {
	id                int
	slug, title, body string
	author            User
	date              mysql.NullTime
}

func (p *Post) Fetch() error {
	var row *sql.Row
	author := User{}

	switch {
	case p.id != 0:
		row = db.QueryRow("SELECT * FROM posts WHERE idPost=?", p.id)
	case p.slug != "":
		row = db.QueryRow("SELECT * FROM posts WHERE slug=?", p.slug)
	default:
		return errors.New("Must provide ID or slug for fetching")
	}

	err := row.Scan(&p.id, &p.slug, &p.title, &p.body, &author.id, &p.date)
	if err == sql.ErrNoRows || err != nil {
		return errors.New("Post not found")
	}
	err = author.Fetch()
	p.author = author
	return err
}
