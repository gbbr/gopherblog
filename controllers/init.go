package controller

import "html/template"

var tpl *template.Template

func init() {
	templates := []string{
		"views/login.html",
		"views/404.html",
		"views/edit.html",
		"views/editPost.html",
		"views/header.html",
		"views/footer.html",
		"views/home.html",
	}

	tpl = template.Must(template.ParseFiles(templates...))
}
