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
