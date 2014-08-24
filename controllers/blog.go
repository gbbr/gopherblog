package controller

import (
	"github.com/gbbr/gopherblog/models"
	"net/http"
)

// Fetches a new post and displays it
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

// Displays home page
func Home(w http.ResponseWriter, r *http.Request) {
	posts, err := models.Posts(200)

	if err != nil || len(r.URL.Path) > 1 {
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
