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
	Draft             bool
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
		err = rows.Scan(&p.Id, &p.Title, &p.Slug, date, &p.Draft)
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

// Saves a post to the database. If it has a set ID it will try to
// update an already existing post, otherwise it will insert a new post
// and generate an ID for it
func (p *Post) Save() error {
	var err error

	if p.Id == 0 {
		_, err = db.Exec(SQL_INSERT_POST, p.Slug, p.Title, p.Body, p.Author.Id, p.Draft)
	} else {
		_, err = db.Exec(SQL_UPDATE_POST, p.Slug, p.Title, p.Body, p.Author.Id, p.Draft, p.Id)
	}

	return err
}

// Scans a fetched row and updates the structure
func (p *Post) update(data *sql.Row) error {
	date := new(mysql.NullTime)
	author := new(User)

	err := data.Scan(&p.Slug, &p.Title, &p.Body, date, &author.Id, &author.Name, &author.Email, &p.Draft)
	if err == sql.ErrNoRows || err != nil {
		return errors.New("Post not found")
	}

	p.Author = *author
	if date.Valid {
		p.Date = date.Time
	}

	return nil
}
