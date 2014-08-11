package controller

import (
	"github.com/gbbr/gopherblog/models"
	"net/http"
	"strconv"
	"strings"
)

// Create a new post
func NewPost(w http.ResponseWriter, r *http.Request, u *models.User) {
	switch r.Method {
	case "POST":
		keys := r.FormValue
		post := &models.Post{
			Slug:   keys("slug"),
			Title:  keys("title"),
			Body:   strings.Trim(keys("body"), " \t\r\n"),
			Author: *u,
			Draft:  keys("draft") != "on",
		}

		post.Save() //todo: handle err
		http.Redirect(w, r, "/manage", http.StatusFound)

	case "GET":
		tpl.ExecuteTemplate(w, "editPost", nil)
	}
}

// View all posts by that user with links to editing them
func Manage(w http.ResponseWriter, r *http.Request, u *models.User) {
	if u.Id == 0 {
		http.Redirect(w, r, "/login?return="+r.URL.Path, http.StatusFound)
	}

	up, _ := models.PostsByUser(u) //todo: handle err
	tpl.ExecuteTemplate(w, "manage", up)
}

// Edit a post
func EditPost(w http.ResponseWriter, r *http.Request, u *models.User) {
	pId, err := strconv.Atoi(r.URL.Path[len("/edit/"):])
	if err != nil {
		tpl.ExecuteTemplate(w, "404", nil)
		return
	}

	switch r.Method {
	case "POST":
		keys := r.FormValue
		post := &models.Post{
			Id:     pId,
			Slug:   keys("slug"),
			Title:  keys("title"),
			Body:   strings.Trim(keys("body"), " \t\r\n"),
			Author: *u,
			Draft:  keys("draft") != "on",
		}

		post.Save() //todo: handle err
		http.Redirect(w, r, "/manage", http.StatusFound)

	case "GET":
		post := &models.Post{Id: pId}

		err = post.Fetch()
		if err != nil || post.Author.Id != u.Id {
			tpl.ExecuteTemplate(w, "404", nil)
			return
		}

		tpl.ExecuteTemplate(w, "editPost", post)
	}
}
