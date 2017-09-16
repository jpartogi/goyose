package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jpartogi/goyose/config"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r := config.Routes()

    log.Fatal(http.ListenAndServe(":"+ port, r))

}