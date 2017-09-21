package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jpartogi/goyose/config"
	"github.com/jpartogi/goyose/view"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" 
	}

	config.Load("config"+string(os.PathSeparator)+"config.json", configuration)
	view.Configure(configuration.View)

	r := config.Routes()

    log.Fatal(http.ListenAndServe(":"+ port, r))

}

var configuration = &config.Configuration{}