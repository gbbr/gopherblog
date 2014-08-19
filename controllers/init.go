package controller

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var (
	// List of files to load templates from
	files = []string{
		"views/login.html",
		"views/404.html",
		"views/manage.html",
		"views/editPost.html",
		"views/header.html",
		"views/footer.html",
		"views/home.html",
		"views/post.html",
	}

	// If flag is set, templates are reloaded on every refresh
	noCache = flag.Bool("nocache", false, "Determines whether template caching should occur.")
	// Global template engine
	tpl *BlogTemplates
)

// Custom template wrapper
type BlogTemplates struct {
	files    []string           // List of files
	compiled *template.Template // Compiled templates
}

// ExecuteTemplate wrapper, if -nocache flag is set, all templates are loaded on every request
func (t *BlogTemplates) ExecuteTemplate(w http.ResponseWriter, name string, data interface{}) error {
	if *noCache {
		t.compiled = template.Must(template.ParseFiles(t.files...))
		log.Printf("[%s] Recompiling templates.", name)
	}

	return t.compiled.ExecuteTemplate(w, name, data)
}

func init() {
	tpl = &BlogTemplates{
		files:    files,
		compiled: template.Must(template.ParseFiles(files...)),
	}
}
