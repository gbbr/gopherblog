package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type httpHandler func(http.ResponseWriter, *http.Request)

// Confirms that a user is authenticated before procdeeding to the
// given HTTP handler, otherwise redirects to login page
func authenticate(dest httpHandler) httpHandler {
	return httpHandler(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("auth")
		if err != nil || err == http.ErrNoCookie {
			http.Redirect(w, r, "/login?return="+r.URL.Path, 307)
			return
		}

		parts := strings.Split(c.Value, ":")

		_, err = strconv.Atoi(parts[0])
		if err != nil {
			http.Redirect(w, r, "/login?return="+r.URL.Path, 307)
			return
		}

		origin := []byte(r.RemoteAddr + r.UserAgent())
		hash := fmt.Sprintf("%x", sha256.Sum256(origin))

		if parts[1] != hash {
			http.Redirect(w, r, "/login?return="+r.URL.Path, 307)
			return
		}

		dest(w, r)
	})
}
