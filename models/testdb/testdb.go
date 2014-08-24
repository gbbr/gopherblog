/*
Package testdb assists with creating and managing
a database destined for testing.

Usage

To use this package, call SetUp from within your test and potentially set up
the SchemaFile path if it differs from the default one. Example set-up with custom
schema file:

	dbtest.Config.SchemaFile = "testsql/schema.sql"
	dbtest.SetUp()

	models.ConnectDb(dbtest.Config.DbString)
*/
package testdb

import (
	"log"
	"os"
	"os/exec"
	"sync"
)

var (
	// Test configuration, db connections, schema files
	Config = struct {
		SchemaFile string
		DbString   string
		user, pass string
	}{
		SchemaFile: "testdb/schema.sql",
		user:       "root",
		pass:       "",
		DbString:   "root@tcp(localhost:3306)/blog_test",
	}

	// Sync helper to run function only once
	once sync.Once
)

// Sets up a test database with data found in the defined schema file.
// Configuration can be altered by changing the Config object.
func SetUp() {
	once.Do(setUpTestDB)
}

func setUpTestDB() {
	file, err := os.Open(Config.SchemaFile)
	if err != nil {
		log.Fatal("Error opening file")
	}
	defer file.Close()

	cmd := exec.Command("mysql", "--user="+Config.user, "--password="+Config.pass)
	cmd.Stdin = file
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
