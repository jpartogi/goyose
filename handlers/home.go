package handlers

import (
	"net/http"

	"github.com/jpartogi/goyose/view"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	v:= view.New("home.html")
    v.Render(w)
}