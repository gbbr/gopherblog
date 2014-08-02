package controller

import (
	"fmt"
	"github.com/backslashed/gopherblog/models"
	"net/http"
	"strconv"
	"strings"
)

// Extracts user's ID from cookie
func getUid(r *http.Request) (int, error) {
	c, err := r.Cookie("auth")
	if err != nil || err == http.ErrNoCookie {
		return 0, err
	}

	uid, err := strconv.Atoi(strings.Split(c.Value, ":")[0])
	if err != nil {
		return 0, nil
	}

	return uid, nil
}

func ViewEdit(w http.ResponseWriter, r *http.Request) {
	uid, err := getUid(r)
	if err != nil {
		http.Redirect(w, r, "/login?return="+r.URL.Path, 307)
	}

	up, err := models.PostsByUser(uid)
	fmt.Fprintf(w, "%+v\n\nEDIT", up)
}

func ViewEditPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%+v\n\nEDIT POST", r)
}
