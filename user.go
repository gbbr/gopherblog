package main

import (
	"database/sql"
	"errors"
)

// Contains information about a user
type User struct {
	id          int
	name, email string
}

// Fetches a user by ID and updates the structure
func (u *User) Fetch() error {
	if u.id == 0 {
		return errors.New("Need ID to fetch")
	}
	row := db.QueryRow("SELECT name, email FROM users WHERE idUser=?", u.id)
	if err := u.update(row); err != nil {
		return errors.New("Error scanning")
	}
	return nil
}

// Updates user structure with data from database
func (u *User) update(row *sql.Row) error {
	err := row.Scan(&u.name, &u.email)
	if err == sql.ErrNoRows || err != nil {
		return errors.New("Could not fetch author")
	}
	return nil
}
