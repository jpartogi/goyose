package tests

import (
	"net/http"
	"testing"

  "github.com/jpartogi/goyose/handlers"
)

func TestPing(t *testing.T) {
  req, rec := TestAPI("/ping", http.MethodGet)  

  handlers.PingHandler(rec, req)

  t.Log(rec.Body.String())

  if rec.Code != http.StatusOK {
  	t.Error("Got ", rec.Code)
  }

  expected:= `{"alive": true}`
  if rec.Body.String() !=  expected{
    t.Error("Got ", rec.Body.String())
  }

}