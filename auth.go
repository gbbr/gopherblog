package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/backslashed/gopherblog/models"
	"net/http"
	"strconv"
	"strings"
)

type httpHandler func(http.ResponseWriter, *http.Request)

// Confirms that a user is authenticated before procdeeding to the
// given HTTP handler, otherwise redirects to login page
func authenticate(dest httpHandler) httpHandler {
	return httpHandler(func(w http.ResponseWriter, r *http.Request) {
		// Does authentication cookie exist?
		c, err := r.Cookie("auth")
		if err != nil || err == http.ErrNoCookie {
			http.Redirect(w, r, "/login?return="+r.URL.Path, 307)
			return
		}

		parts := strings.Split(c.Value, ":")
		// Can we extract an integer?
		uid, err := strconv.Atoi(parts[0])
		if err != nil {
			http.Redirect(w, r, "/login?return="+r.URL.Path, 307)
			return
		}

		// Is there a user with that ID?
		user := &models.User{Id: uid}
		if err := user.Fetch(); err != nil {
			http.Redirect(w, r, "/login?return="+r.URL.Path, 307)
			return
		}

		origin := []byte(string(user.Email) + r.RemoteAddr + r.UserAgent())
		hash := fmt.Sprintf("%x", sha256.Sum256(origin))

		// Does the hash match the origin?
		if parts[1] != hash {
			http.Redirect(w, r, "/login?return="+r.URL.Path, 307)
			return
		}

		dest(w, r)
	})
}
