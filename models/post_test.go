package models

import (
	"reflect"
	"testing"
	"time"

	"github.com/gbbr/gopherblog/models/testdb"
)

func TestPostsWithLimit(t *testing.T) {
	testdb.SetUp()

	ConnectDb(testdb.Config.DbString)
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
	testdb.SetUp()

	ConnectDb(testdb.Config.DbString)
	defer CloseDb()

	posts, err := PostsByUser(&User{Id: 1})
	if posts[0].Draft != true || posts[4].Slug != "slug-one" ||
		len(posts) != 5 || err != nil {

		t.Log("Unexpected result")
		t.Fail()
	}
}

func TestPostFetch(t *testing.T) {
	testdb.SetUp()

	ConnectDb(testdb.Config.DbString)
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

func TestPostSaveNew(t *testing.T) {
	testdb.SetUp()

	ConnectDb(testdb.Config.DbString)
	defer CloseDb()

	timeCompromise := time.Now()
	want := &Post{
		Title: "My shiny new post",
		Slug:  "shiny-post",
		Body:  "Look at the bling in this post",
		Author: User{
			Id:    1,
			Name:  "Jeremy",
			Email: "jeremy@email.com",
		},
		Date:  timeCompromise,
		Draft: false,
		Tags:  []string{"A", "B", "C"},
	}

	err := want.Save()
	if err != nil {
		t.Log("Error saving post")
		t.Fail()
	}

	post := &Post{Id: want.Id}
	post.Fetch()

	post.Date = timeCompromise // Ignore time on deep equals
	if !reflect.DeepEqual(post, want) {
		t.Log("Did not save post correctly")
		t.Fail()
	}

	post.Delete()
}

func TestPostUpdate(t *testing.T) {
	testdb.SetUp()

	ConnectDb(testdb.Config.DbString)
	defer CloseDb()

	original := &Post{Id: 11}
	original.Fetch()

	timeCompromise := time.Now()
	want := &Post{
		Id:    11,
		Title: "My shiny new post",
		Slug:  "shiny-post",
		Body:  "Look at the bling in this post",
		Author: User{
			Id:    1,
			Name:  "Jeremy",
			Email: "jeremy@email.com",
		},
		Date:  timeCompromise,
		Draft: false,
		Tags:  []string{"A", "B", "C"},
	}

	err := want.Save()
	if err != nil {
		t.Log("Error saving post")
		t.Fail()
	}

	post := &Post{Id: want.Id}
	post.Fetch()

	post.Date = timeCompromise // Ignore time on deep equals
	if !reflect.DeepEqual(post, want) {
		t.Log("Did not save post correctly")
		t.Fail()
	}

	err = original.Save()
	if err != nil {
		t.Log("Error restoring original")
	}
}

func TestPostDelete(t *testing.T) {
	testdb.SetUp()

	ConnectDb(testdb.Config.DbString)
	defer CloseDb()

	original := &Post{Id: 11}
	original.Fetch()

	clone := new(Post)
	*clone = *original
	
	err := clone.Delete()
	if err != nil {
		t.Fatal("Error deleting post")
	}

	if clone.Id != 0 {
		t.Log("Did not delete post")
		t.Fail()
	}

	original.Save()
}
