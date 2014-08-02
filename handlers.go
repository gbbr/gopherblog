package main

import (
	"fmt"
	"net/http"
)

func viewPost(w http.ResponseWriter, r *http.Request) {
	post := &Post{
		slug: r.URL.Path[len("/posts/")-1:],
	}

	err := post.Fetch()
	if err != nil {
		fmt.Fprintf(w, "%+v", "Post not found")
	}

	fmt.Fprintf(w, "%+v", post)
}
