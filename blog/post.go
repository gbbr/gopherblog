package blog

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Holds post data
type Post struct {
	id          int32
	title, body string
	author      User
	date        time.Time
}

func (p *Post) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.URL.Path[6:], 10, 32)
	if err != nil {
		fmt.Fprint(w, "Invalid ID: ", req.URL.Path[5:])
	}
	fmt.Fprintf(w, "%d", id)
}
