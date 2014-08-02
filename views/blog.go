package views

import (
	"fmt"
	"github.com/backslashed/gopherblog/models"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	templates := []string{
		"views/templates/login.html",
		"views/templates/404.html",
	}

	tpl = template.Must(template.ParseFiles(templates...))
}

// Fetches a new post and displays it
func ViewPost(w http.ResponseWriter, r *http.Request) {
	post := &models.Post{
		Slug: r.URL.Path[len("/posts/")-1:],
	}

	err := post.Fetch()
	if err != nil {
		tpl.ExecuteTemplate(w, "404.html", nil)
		return
	}

	fmt.Fprintf(w, "%+v", post)
}

// Displays home page
func ViewHome(w http.ResponseWriter, r *http.Request) {
	posts, err := models.Posts(200)

	if err != nil || len(r.URL.Path) > 1 {
		tpl.ExecuteTemplate(w, "404.html", nil)
		return
	}

	for _, post := range posts {
		fmt.Fprintf(w, "%+v\n\n", post)
	}
}
