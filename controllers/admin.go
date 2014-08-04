package controller

import (
	"github.com/backslashed/gopherblog/models"
	"net/http"
	"strconv"
)

func ViewEdit(w http.ResponseWriter, r *http.Request, u *models.User) {
	if u.Id == 0 {
		http.Redirect(w, r, "/login?return="+r.URL.Path, 307)
	}

	up, _ := models.PostsByUser(u) //todo: handle err
	tpl.ExecuteTemplate(w, "edit", up)
}

func ViewEditPost(w http.ResponseWriter, r *http.Request, u *models.User) {
	pId, err := strconv.Atoi(r.URL.Path[len("/edit/"):])
	if err != nil {
		tpl.ExecuteTemplate(w, "404", nil)
		return
	}

	post := &models.Post{Id: pId}
	err = post.Fetch()
	if err != nil {
		tpl.ExecuteTemplate(w, "404", nil)
		return
	}

	tpl.ExecuteTemplate(w, "editPost", post)
}
