package handlers

import (
	"net/http"
	"html/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/home.html")  // Parse template file.
    t.Execute(w, nil)
}