package models

import (
	"database/sql"
	"errors"
	mysql "github.com/go-sql-driver/mysql"
	"time"
)

const (
	SQL_POST_BY_ID   = "SELECT title, body, idUser, date FROM posts WHERE idPost=?"
	SQL_POST_BY_SLUG = "SELECT title, body, idUser, date FROM posts WHERE slug=?"
)

type Post struct {
	Id                int
	Slug, Title, Body string
	Author            User
	Date              time.Time
}

// Fetches data from database by ID or slug (whichever one is available)
// and updates structure
func (p *Post) Fetch() error {
	var data *sql.Row

	switch {
	case p.Id != 0:
		data = db.QueryRow(SQL_POST_BY_ID, p.Id)
	case p.Slug != "":
		data = db.QueryRow(SQL_POST_BY_SLUG, p.Slug)
	default:
		return errors.New("Must provide ID or slug for fetching")
	}

	err := p.update(data)
	if err != nil {
		return errors.New("Error scanning row")
	}

	return nil
}

// Scans a fetched row and updates the structure
func (p *Post) update(data *sql.Row) error {
	date := new(mysql.NullTime)
	author := new(User)

	err := data.Scan(&p.Title, &p.Body, &author.Id, date)
	if err == sql.ErrNoRows || err != nil {
		return errors.New("Post not found")
	}

	if date.Valid {
		p.Date = date.Time
	}

	err = author.Fetch()
	if err != nil {
		return err
	}

	p.Author = *author
	return nil
}
