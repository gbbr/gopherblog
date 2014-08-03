package controller

import (
	"fmt"
	"github.com/backslashed/gopherblog/models"
	"net/http"
)

func ViewEdit(w http.ResponseWriter, r *http.Request, u *models.User) {
	if u.Id == 0 {
		http.Redirect(w, r, "/login?return="+r.URL.Path, 307)
	}

	up, _ := models.PostsByUser(u) //todo: handle err
	fmt.Fprintf(w, "%+v\n\nEDIT", up)
}

func ViewEditPost(w http.ResponseWriter, r *http.Request, u *models.User) {
	fmt.Fprintf(w, "%+v\n\nEDIT POST", r)
}
