package controller

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

// If flag is set, templates are reloaded on every refresh
var noCache = flag.Bool("nocache", false, "Determines whether template caching should occur.")

// Global template engine
var tpl *BlogTemplates

// Custom template wrapper
type BlogTemplates struct {
	files    []string           // List of files
	compiled *template.Template // Compiled templates
}

// ExecuteTemplate wrapper, if -nocache flag is set, all templates are loaded on every request
func (t *BlogTemplates) ExecuteTemplate(w http.ResponseWriter, name string, data interface{}) error {
	if *noCache {
		t.compiled = template.Must(template.ParseFiles(t.files...))
		log.Println("Recompiling templates.")
	}

	return t.compiled.ExecuteTemplate(w, name, data)
}

func init() {
	if *noCache {
		log.Println("Running with no-cache...")
	}

	tpl = &BlogTemplates{
		files: []string{
			"views/login.html",
			"views/404.html",
			"views/edit.html",
			"views/editPost.html",
			"views/header.html",
			"views/footer.html",
			"views/home.html",
		},
		compiled: template.Must(template.ParseFiles(templates...)),
	}
}
