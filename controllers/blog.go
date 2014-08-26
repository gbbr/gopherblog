package controller

import (
	"net/http"
	"strings"

	"github.com/gbbr/gopherblog/models"
)

// Displays all posts or filters them by tag according to the URL path.
// If the path contains "/tag/:tag" it will filter by tag, otherwise it
// will show the homepage. If there are trailing characters after the URL
// it will show 404 page.
func Posts(w http.ResponseWriter, r *http.Request) {
	var (
		posts []models.Post
		err   error
	)

	if strings.HasPrefix(r.URL.Path, "/tag/") {
		posts, err = models.PostsByTag(r.URL.Path[len("/tag/"):])
	} else {
		posts, err = models.Posts(200)
	}

	if err != nil {
		tpl.ExecuteTemplate(w, "404", nil)
		return
	}

	tags, err := models.TagsAll()
	if err != nil {
		tpl.ExecuteTemplate(w, "404", nil)
		return
	}

	tpl.ExecuteTemplate(w, "home", struct {
		Posts []models.Post
		Tags  []string
	}{
		posts,
		tags,
	})
}

// Fetches a new post and displays it. Goes to 404 Not Found
// page if post is invalid or URL is odd
func Post(w http.ResponseWriter, r *http.Request) {
	post := &models.Post{
		Slug: r.URL.Path[len("/post/"):],
	}

	err := post.Fetch()
	if err != nil {
		tpl.ExecuteTemplate(w, "404", nil)
		return
	}

	tpl.ExecuteTemplate(w, "post", post)
}
