package controller

import (
	"fmt"
	"net/http"
)

func ViewEdit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%+v\n\nEDIT", r)
}

func ViewEditPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%+v\n\nEDIT POST", r)
}
