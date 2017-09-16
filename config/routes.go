package config

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jpartogi/goyose/handlers"

)

func Routes() *mux.Router {
	r := mux.NewRouter()

	// routes goes here
    r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/ping", handlers.PingHandler).Methods("GET")


    // static files
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	
    return r
}