package models

import (
	"database/sql"
	"errors"
)

const SQL_USER_BYID = "SELECT name, email FROM users WHERE idUser=?"

// Contains information about a user
type User struct {
	Id          int
	Name, Email string
}

// Fetches a user by ID and updates the structure
func (u *User) Fetch() error {
	if u.Id == 0 {
		return errors.New("Need ID to fetch")
	}
	row := db.QueryRow(SQL_USER_BYID, u.Id)

	err := u.update(row)
	if err != nil {
		return errors.New("Error scanning")
	}

	return nil
}

// Updates user structure with data from database
func (u *User) update(row *sql.Row) error {
	err := row.Scan(&u.Name, &u.Email)
	if err == sql.ErrNoRows || err != nil {
		return errors.New("Could not fetch author")
	}

	return nil
}
