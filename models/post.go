package models

import (
	"database/sql"
	"errors"
	mysql "github.com/go-sql-driver/mysql"
	"time"
)

type Post struct {
	Id                int
	Slug, Title, Body string
	Author            User
	Date              time.Time
}

const (
	SQL_POST_BY_ID = `SELECT title, body, date, idUser, users.name, users.email 
		FROM posts 
		INNER JOIN users USING(idUser)
		WHERE idPost=?`
	SQL_POST_BY_SLUG = `SELECT title, body, date, idUser, users.name, users.email 
		FROM posts 
		INNER JOIN users USING(idUser)
		WHERE slug=?`
)

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

	err := data.Scan(&p.Title, &p.Body, date, &author.Id, &author.Name, &author.Email)
	if err == sql.ErrNoRows || err != nil {
		return errors.New("Post not found")
	}

	p.Author = *author
	if date.Valid {
		p.Date = date.Time
	}

	return nil
}
