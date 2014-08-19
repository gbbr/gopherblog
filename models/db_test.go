package models

import (
	"log"
	"os"
	"os/exec"
	"sync"
)

var (
	// Test configuration, db connections, schema files
	testConfig = struct {
		schemaFile string
		user, pass string
		dbString   string
	}{
		schemaFile: "../sql/schema_test.sql",
		user:       "root",
		pass:       "root",
		dbString:   "root:root@tcp(localhost:3306)/blog_test",
	}

	// Sync helper to run function only once
	once sync.Once
)

// Sets up test DB from schema file
func setUp() {
	file, err := os.Open(testConfig.schemaFile)
	if err != nil {
		log.Fatal("Error opening file")
	}
	defer file.Close()

	cmd := exec.Command("mysql", "--user="+testConfig.user, "--password="+testConfig.pass)
	cmd.Stdin = file
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
