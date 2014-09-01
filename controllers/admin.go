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
		err := savePost(0, r.FormValue, u)
		if err != nil {
			tpl.ExecuteTemplate(w, "404", nil)
			return
		}

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

	up, err := models.PostsByUser(u)
	if err != nil {
		tpl.ExecuteTemplate(w, "404", nil)
		return
	}

	tpl.ExecuteTemplate(w, "manage", up)
}

// Edit a post
func EditPost(w http.ResponseWriter, r *http.Request, u *models.User) {
	pId, err := strconv.Atoi(r.URL.Path[len("/edit/"):])
	if err != nil {
		tpl.ExecuteTemplate(w, "404", nil)
		return
	}

	// If the request method is POST, save any changes to
	// the post
	if r.Method == "POST" {
		err = savePost(pId, r.FormValue, u)
		if err != nil {
			tpl.ExecuteTemplate(w, "404", nil)
			return
		}
	}

	post := &models.Post{Id: pId}

	err = post.Fetch()
	if err != nil || post.Author.Id != u.Id {
		tpl.ExecuteTemplate(w, "404", nil)
		return
	}

	tpl.ExecuteTemplate(w, "editPost", post)
}

// Saves a post extracted from a form value function
// if ID is 0 the post will be inserted
func savePost(id int, keys func(string) string, u *models.User) error {
	var tags []string

	for _, tag := range strings.Split(keys("tags"), ",") {
		tags = append(tags, strings.Trim(tag, " "))
	}

	post := &models.Post{
		Id:       id,
		Slug:     keys("slug"),
		Title:    keys("title"),
		Abstract: keys("abstract"),
		Body:     strings.Trim(keys("body"), " \t\r\n"),
		Author:   *u,
		Draft:    keys("draft") != "on",
		Tags:     tags,
	}

	return post.Save()
}
