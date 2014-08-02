package views

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/backslashed/gopherblog/models"
	"net/http"
)

func ViewLogin(w http.ResponseWriter, r *http.Request) {
	tplData := struct{ Msg, ReturnUrl string }{
		ReturnUrl: r.URL.Query().Get("return"),
	}

	if r.Method == "POST" {
		md5pass := fmt.Sprintf("%x", md5.Sum([]byte(r.FormValue("password"))))

		user := &models.User{
			Email:    r.FormValue("login"),
			Password: md5pass,
		}

		if user.LoginCorrect() {
			redirect := r.FormValue("redirectUrl")
			if len(redirect) == 0 {
				redirect = "/"
			}

			origin := []byte(r.RemoteAddr + r.UserAgent())
			val := fmt.Sprintf("%d:%x", user.Id, sha256.Sum256(origin))
			http.SetCookie(w, &http.Cookie{
				Name:  "auth",
				Value: val,
			})

			http.Redirect(w, r, redirect, 307)
		}

		tplData.Msg = "Invalid login and password combination."
	}

	tpl.ExecuteTemplate(w, "login.html", tplData)
}

func ViewEdit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%+v\n\nEDIT", r)
}

func ViewEditPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%+v\n\nEDIT POST", r)
}