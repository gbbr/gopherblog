package blog

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

// Holds the blog's configuration
type Config struct {
	Host, Port string
	DbString   string
}

// Database connection
var db *sql.DB

// Creates a new blog instance
func Start(conf Config) error {
	db, err := sql.Open("mysql", conf.DbString)
	if err != nil {
		errors.New("Error opening SQL connnection")
	}
	if err := db.Ping(); err != nil {
		errors.New("Error connecting to database")
	}
	return nil
}

// Returns an active database connection or an error
func Db() (*sql.DB, error) {
	if db == nil {
		return db, errors.New("Inactive database connection")
	}
	return db, nil
}
