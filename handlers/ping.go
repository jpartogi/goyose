package handlers

import (
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	ping:= []byte(`{"alive": true}`)

	w.Header().Set("Content-Type", "application/json")
	w.Write(ping)
}