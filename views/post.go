package views

import (
	"fmt"
	"github.com/backslashed/gopherblog/models"
	"net/http"
)

// Fetches a new post and displays it
func ViewPost(w http.ResponseWriter, r *http.Request) {
	post := &models.Post{
		Slug: r.URL.Path[len("/posts/")-1:],
	}

	err := post.Fetch()
	if err != nil {
		fmt.Fprintf(w, "%+v", "Post not found")
	}

	fmt.Fprintf(w, "%+v", post)
}

// Displays home page
func ViewHome(w http.ResponseWriter, r *http.Request) {
	posts, err := models.Posts(200)

	if err != nil {
		fmt.Fprint(w, "404")
	}

	for _, post := range posts {
		fmt.Fprintf(w, "%+v\n\n", post)
	}
}
