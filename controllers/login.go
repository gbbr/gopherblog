package controller

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/backslashed/gopherblog/models"
	"net/http"
)

// Displays the login form template. Interprets both GET and POST
func ViewLogin(w http.ResponseWriter, r *http.Request) {
	tplData := struct{ Msg, ReturnUrl string }{
		ReturnUrl: r.URL.Query().Get("return"),
	}

	if r.Method == "POST" {
		validateLoginForm(w, r)
		tplData.Msg = "Invalid login and password combination."
	}

	tpl.ExecuteTemplate(w, "login", tplData)
}

// Validates the login form's POST and redirects the user if login
// is correct
func validateLoginForm(w http.ResponseWriter, r *http.Request) {
	md5pass := fmt.Sprintf("%x", md5.Sum([]byte(r.FormValue("password"))))

	user := &models.User{
		Email:    r.FormValue("login"),
		Password: md5pass,
	}

	if user.LoginCorrect() {
		origin := []byte(string(user.Email) + r.RemoteAddr + r.UserAgent())
		val := fmt.Sprintf("%d:%x", user.Id, sha256.Sum256(origin))
		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: val,
		})

		w.Header().Set("Method", "GET")
		http.Redirect(w, r, r.FormValue("redirectUrl"), 307)
	}
}
