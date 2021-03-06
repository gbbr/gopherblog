package models

import (
	"database/sql"
	"errors"
	"fmt"
	mysql "github.com/go-sql-driver/mysql"
	"html/template"
	"strings"
	"time"
)

type Post struct {
	Id          int
	Slug        string
	Title, Body string
	Abstract    string
	Author      User
	Date        time.Time
	Draft       bool
	Tags        []string
}

// Fetches number of posts from the database ordered by date
func Posts(limit int) (posts []Post, err error) {
	rows, err := db.Query(SQL_ALL_POSTS, limit)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		post, author, date := new(Post), new(User), new(mysql.NullTime)
		err = rows.Scan(&post.Slug, &post.Title, &post.Abstract, date, &author.Id, &author.Name)
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
	if err != nil {
		return
	}
	defer rows.Close()

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

// Retrieves all posts that match a given tag
// and are not drafts, ordered by date
func PostsByTag(tag string) (posts []Post, err error) {
	if len(tag) == 0 {
		return nil, errors.New("Invalid tag")
	}

	rows, err := db.Query(SQL_POSTS_BY_TAG, tag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	p := new(Post)
	d := new(mysql.NullTime)

	for rows.Next() {
		err := rows.Scan(&p.Slug, &p.Title, &p.Abstract, d, &p.Author.Id, &p.Author.Name)
		if err != nil {
			return nil, err
		}

		if d.Valid {
			p.Date = d.Time
		} else {
			return nil, err
		}

		posts = append(posts, *p)
	}

	return
}

// Returns a list of unique tags
func TagsAll() ([]string, error) {
	rows, err := db.Query(SQL_ALL_TAGS)
	if err != nil {
		return nil, err
	}

	var (
		tags []string
		tag  string
	)

	for rows.Next() {
		err := rows.Scan(&tag)
		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}

// Fetches one post from the database based on ID
// or slug
func (p *Post) Fetch() error {
	var (
		rows *sql.Rows
		err  error
	)

	switch {
	case p.Id != 0:
		rows, err = db.Query(SQL_POST_BY_ID, p.Id)
	case p.Slug != "":
		rows, err = db.Query(SQL_POST_BY_SLUG, p.Slug)
	default:
		return errors.New("Must provide ID or slug for fetching")
	}

	if err != nil {
		return err
	}

	tag := new(sql.NullString)
	date := new(mysql.NullTime)

	for rows.Next() {
		err := rows.Scan(&p.Id, &p.Slug, &p.Title, &p.Abstract,
			&p.Body, date, &p.Author.Id, &p.Author.Name,
			&p.Author.Email, &p.Draft, tag)

		if err != nil {
			return err
		}

		if date.Valid {
			p.Date = date.Time
		}

		if tag.Valid {
			p.Tags = append(p.Tags, tag.String)
		} else {
			break
		}
	}

	rows.Close()
	return nil
}

// Saves a post to the database. If it has a set ID it will try to
// update an already existing post, otherwise it will insert a new post
// and generate an ID for it. An unset ID is an ID of value 0.
func (p *Post) Save() error {
	var (
		err    error
		result sql.Result
	)

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if p.Id == 0 {
		result, err = tx.Exec(SQL_INSERT_POST, p.Slug, p.Title, p.Abstract, p.Body, p.Author.Id, p.Draft)
		if err != nil {
			tx.Rollback()
			return err
		}

		var id64 int64
		id64, err = result.LastInsertId()
		p.Id = int(id64)
	} else {
		result, err = tx.Exec(SQL_UPDATE_POST, p.Slug, p.Title, p.Abstract, p.Body, p.Author.Id, p.Draft, p.Id)
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(SQL_REMOVE_TAGS, p.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if len(p.Tags) > 0 {
		stmt, err := tx.Prepare(SQL_INSERT_TAGS)
		if err != nil {
			tx.Rollback()
			return err
		}

		for _, tag := range p.Tags {
			if len(strings.Trim(tag, " ")) > 0 {
				_, err = stmt.Exec(p.Id, tag)
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}

	tx.Commit()
	return nil
}

// Deletes the post
func (p *Post) Delete() error {
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete post
	_, err = tx.Exec(SQL_DELETE_POST, p.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Remove tags
	_, err = tx.Exec(SQL_REMOVE_TAGS, p.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Attempt to commit changes
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	*p = Post{}

	return err
}

// Template helper functions

func (p *Post) FormattedDate() string {
	year, month, day := p.Date.Date()
	return fmt.Sprintf("%d %s %d", day, month, year)
}

func (p *Post) BodyHTML() template.HTML { return template.HTML(p.Body) }
func (p *Post) TagsString() string      { return strings.Join(p.Tags, ", ") }
