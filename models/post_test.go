package models

import (
	"log"
	"os"
	"os/exec"
	"sync"
	"testing"
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

func TestPostsWithLimit(t *testing.T) {
	once.Do(setUp)

	ConnectDb(testConfig.dbString)
	defer CloseDb()

	// Test retrieval of all (considering there's under 100)
	posts, err := Posts(100)
	if len(posts) != 7 || err != nil {
		t.Log("Failed to retrieve all")
		t.Fail()
	}

	// Test limited retrieval
	posts, err = Posts(3)
	if len(posts) != 3 || err != nil {
		t.Log("Failed to retrieve limited")
		t.Fail()
	}
}

func TestPostsByUser(t *testing.T) {
	once.Do(setUp)

	ConnectDb(testConfig.dbString)
	defer CloseDb()

	posts, err := PostsByUser(&User{Id: 1})
	if posts[0].Draft != true || posts[4].Slug != "slug-one" ||
		len(posts) != 5 || err != nil {

		t.Log("Unexpected result")
		t.Fail()
	}
}

func TestFetch(t *testing.T) {
	once.Do(setUp)

	ConnectDb(testConfig.dbString)
	defer CloseDb()

	// Test fetch by slug
	post := &Post{Slug: "slug-two"}
	if err := post.Fetch(); err != nil {
		t.Log("An error occurred while fetching post")
		t.Fatal()
	}

	if post.Id != 12 || post.Title != "Title Two" {
		t.Log("Did not retrieve correct post")
		t.Fail()
	}

	// Test fetch by ID
	post = &Post{Id: 20}
	if err := post.Fetch(); err != nil {
		t.Log("An error occurred while fetching post")
		t.Fatal()
	}

	if post.Id != 20 || post.Title != "My Post Four" || post.Slug != "mypost-four" {
		t.Log("Did not retrieve correct post")
		t.Fail()
	}

	// Test if we fetched the correct author
	if post.Author.Id != 2 || post.Author.Name != "Mathias" || post.Author.Email != "mathias@company.it" {
		t.Log("Didn't fetch the right author")
		t.Fail()
	}
}
