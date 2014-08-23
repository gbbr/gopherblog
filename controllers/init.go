package controller

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"sync"
)

// If flag is set, templates are reloaded on every refresh
var noCache = flag.Bool("nocache", false, "Determines whether template caching should occur.")

// Global template engine
var tpl *BlogTemplate

// Custom template wrapper
type BlogTemplate struct {
	*template.Template
	files []string
	mu    sync.Mutex
}

func (t *BlogTemplate) ExecuteTemplate(w http.ResponseWriter, name string, data interface{}) error {
	if *noCache {
		t.mu.Lock()
		defer t.mu.Unlock()
		log.Printf("[%s] Recompiling templates.", name)
		t.Template = template.Must(template.ParseFiles(t.files...))
	}

	return t.Template.ExecuteTemplate(w, name, data)
}

func init() {
	tpl = &BlogTemplate{
		files: []string{
			"views/login.html",
			"views/404.html",
			"views/manage.html",
			"views/editPost.html",
			"views/header.html",
			"views/footer.html",
			"views/home.html",
			"views/post.html",
		},
	}

	tpl.Template = template.Must(template.ParseFiles(tpl.files...))
}
