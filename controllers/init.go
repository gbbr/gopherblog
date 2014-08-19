package controller

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

// Templates
var templates = []string{
	"views/login.html",
	"views/404.html",
	"views/edit.html",
	"views/editPost.html",
	"views/header.html",
	"views/footer.html",
	"views/home.html",
}

// If flag is set, templates are reloaded on every refresh
var noCache = flag.Bool("nocache", false, "Determines whether template caching should occur.")

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

var tpl *BlogTemplates

func init() {
	if *noCache {
		log.Println("Running with no-cache...")
	}

	tpl = &BlogTemplates{
		files:    templates,
		compiled: template.Must(template.ParseFiles(templates...)),
	}
}
