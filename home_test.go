package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
  "strings"

  "github.com/jpartogi/goyose/handlers"
)

func TestSomething(t *testing.T) {
  
  req := httptest.NewRequest(http.MethodGet, "/", nil)
  req.Header.Set("Content-Type", "text/html")
  rec := httptest.NewRecorder()

  handlers.HomeHandler(rec, req)

  t.Log(rec.Body.String())

  if rec.Code != http.StatusOK {
  	t.Error("Got ", rec.Code)
  }

  if !strings.Contains(rec.Body.String(), "Hello Yose") {
    t.Error("Does not contain expected string")
  }

}