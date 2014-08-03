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

// Fetches number of posts from the database ordered by date
func Posts(limit int) (posts []Post, err error) {
	rows, err := db.Query(SQL_ALL_POSTS, limit)
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		post, author, date := new(Post), new(User), new(mysql.NullTime)
		err = rows.Scan(&post.Slug, &post.Title, date, &author.Id, &author.Name)
		if err != nil {
			return
		}

		post.Author = *author
		if date.Valid {
			post.Date = date.Time
		}

		posts = append(posts, *post)
	}

	err = nil
	return
}

// Fetches all posts by user's ID
func PostsByUser(u *User) (posts []Post, err error) {
	rows, err := db.Query(SQL_POSTS_BY_USER, u.Id)
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		p, date := new(Post), new(mysql.NullTime)
		err = rows.Scan(&p.Title, &p.Slug, date)
		if err != nil {
			return
		}

		if date.Valid {
			p.Date = date.Time
		} else {
			return
		}

		posts = append(posts, *p)
	}

	err = nil
	return
}

// Fetches one post from the database based on ID
// or slug
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
