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
type authHandler func(http.ResponseWriter, *http.Request, *models.User)

// Confirms that a user is authenticated & author before proceeding to the
// given HTTP handler, otherwise redirects to login page
func authenticate(dest authHandler) httpHandler {
	return httpHandler(func(w http.ResponseWriter, r *http.Request) {
		user, isValid := isValidRequest(r)
		if !isValid {
			http.Redirect(w, r, "/login?return="+r.URL.Path, 307)
			return
		}

		dest(w, r, user)
	})
}

// Validates cookie and authentication token
func isValidRequest(r *http.Request) (*models.User, bool) {
	// Does authentication cookie exist?
	c, err := r.Cookie("auth")
	if err != nil || err == http.ErrNoCookie {
		return new(models.User), false
	}

	parts := strings.Split(c.Value, ":")
	// Can we extract an integer?
	uid, err := strconv.Atoi(parts[0])
	if err != nil {
		return new(models.User), false
	}

	// Is there an author with that ID?
	user := &models.User{Id: uid}
	if err := user.Fetch(); err != nil {
		return new(models.User), false
	}

	origin := []byte(string(user.Email) + r.RemoteAddr + r.UserAgent())
	hash := fmt.Sprintf("%x", sha256.Sum256(origin))

	// Does the hash match the origin?
	if parts[1] != hash {
		return new(models.User), false
	}

	return user, true
}
