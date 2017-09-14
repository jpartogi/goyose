package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"	
	"github.com/jpartogi/goyose/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

    r := mux.NewRouter()
    r.HandleFunc("/", handlers.HomeHandler).Methods("GET")

    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

    log.Fatal(http.ListenAndServe(":"+ port, r))

}