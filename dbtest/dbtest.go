package dbtest

import (
	"log"
	"os"
	"os/exec"
	"sync"
)

var (
	// Test configuration, db connections, schema files
	TestConfig = struct {
		SchemaFile string
		user, pass string
		DbString   string
	}{
		SchemaFile: "../dbtest/schema.sql",
		user:       "root",
		pass:       "",
		DbString:   "root@tcp(localhost:3306)/blog_test",
	}

	// Sync helper to run function only once
	once sync.Once
)

func SetUp() {
	once.Do(setUpTestDB)
}

// Sets up test DB from schema file
func setUpTestDB() {
	file, err := os.Open(TestConfig.SchemaFile)
	if err != nil {
		log.Fatal("Error opening file")
	}
	defer file.Close()

	cmd := exec.Command("mysql", "--user="+TestConfig.user, "--password="+TestConfig.pass)
	cmd.Stdin = file
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
