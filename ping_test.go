package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

  "github.com/jpartogi/goyose/handlers"
)

func TestPing(t *testing.T) {
  
  req := httptest.NewRequest(http.MethodGet, "/ping", nil)
  req.Header.Set("Content-Type", "application/json")
  rec := httptest.NewRecorder()

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