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

// Confirms that a user is authenticated & author before proceeding to the
// given HTTP handler, otherwise redirects to login page
func authenticate(dest httpHandler) httpHandler {
	return httpHandler(func(w http.ResponseWriter, r *http.Request) {
		if !isCookieValid(w, r) {
			http.Redirect(w, r, "/login?return="+r.URL.Path, 307)
			return
		}

		dest(w, r)
	})
}

// Validates cookie and authentication token
func isCookieValid(w http.ResponseWriter, r *http.Request) bool {
	// Does authentication cookie exist?
	c, err := r.Cookie("auth")
	if err != nil || err == http.ErrNoCookie {
		return false
	}

	parts := strings.Split(c.Value, ":")
	// Can we extract an integer?
	uid, err := strconv.Atoi(parts[0])
	if err != nil {
		return false
	}

	// Is there an author with that ID?
	user := &models.User{Id: uid}
	if err := user.Fetch(); err != nil {
		return false
	}

	origin := []byte(string(user.Email) + r.RemoteAddr + r.UserAgent())
	hash := fmt.Sprintf("%x", sha256.Sum256(origin))

	// Does the hash match the origin?
	if parts[1] != hash {
		return false
	}

	return true
}
