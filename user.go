package main

import (
	"database/sql"
	"errors"
)

type User struct {
	id          int
	name, email string
	isAuthor    bool
}

func (u *User) Fetch() error {
	if u.id == 0 {
		return errors.New("Need ID to fetch")
	}
	row := db.QueryRow("SELECT name, email, isAuthor FROM users WHERE idUser=?", u.id)
	err := row.Scan(&u.name, &u.email, &u.isAuthor)
	if err == sql.ErrNoRows || err != nil {
		return errors.New("Could not fetch author")
	}
	return nil
}
