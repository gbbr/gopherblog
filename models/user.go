package models

import (
	"database/sql"
	"errors"
)

// Contains information about a user
type User struct {
	Id          int
	Name, Email string
	Password    string
}

// Fetches a user by ID and updates the structure
func (u *User) Fetch() error {
	if u.Id == 0 {
		return errors.New("Need ID to fetch")
	}
	row := db.QueryRow(SQL_USER_BY_ID, u.Id)

	err := u.update(row)
	if err != nil {
		return errors.New("Error scanning")
	}

	return nil
}

// Verifies if a password & e-mail combination is
// correct for the user and fetches rest of data
func (u *User) LoginCorrect() bool {
	if len(u.Email) == 0 || len(u.Password) == 0 {
		return false
	}

	row := db.QueryRow(SQL_USER_AUTH, u.Email, u.Password)
	err := row.Scan(&u.Name, &u.Id)

	if err != nil || err == sql.ErrNoRows {
		return false
	}

	return true
}

// Updates user structure with data from database
func (u *User) update(row *sql.Row) error {
	err := row.Scan(&u.Name, &u.Email)
	if err == sql.ErrNoRows || err != nil {
		return errors.New("Could not fetch author")
	}

	return nil
}
