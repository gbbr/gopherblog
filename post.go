package main

import (
	"database/sql"
	"errors"
	mysql "github.com/go-sql-driver/mysql"
	"time"
)

type Post struct {
	id                int
	slug, title, body string
	author            User
	date              time.Time
}

// Fetches data from database by ID or slug (whichever one is available)
// and updates structure
func (p *Post) Fetch() error {
	var data *sql.Row

	switch {
	case p.id != 0:
		data = p.byId(p.id)
	case p.slug != "":
		data = p.bySlug(p.slug)
	default:
		return errors.New("Must provide ID or slug for fetching")
	}

	if err := p.update(data); err != nil {
		return errors.New("Error scanning row")
	}
	return nil
}

// Query DB by ID
func (p *Post) byId(id int) *sql.Row {
	return db.QueryRow("SELECT title, body, idUser, date FROM posts WHERE idPost=?", id)
}

// Query DB by slug
func (p *Post) bySlug(slug string) *sql.Row {
	return db.QueryRow("SELECT title, body, idUser, date FROM posts WHERE slug=?", slug)
}

// Scans a fetched row and updates the structure
func (p *Post) update(data *sql.Row) error {
	date := new(mysql.NullTime)
	author := new(User)
	err := data.Scan(&p.title, &p.body, &author.id, date)

	if err == sql.ErrNoRows || err != nil {
		return errors.New("Post not found")
	}

	if err := author.Fetch(); err != nil {
		return err
	}

	if date.Valid {
		p.date = date.Time
	}

	p.author = *author
	return nil
}
