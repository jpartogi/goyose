package tests

import (
	"testing"
  "strings"
  "net/http"

  "github.com/jpartogi/goyose/handlers"
)

func TestHome(t *testing.T) {
  req, rec := TestHtmlPage("/", http.MethodGet)

  handlers.HomeHandler(rec, req)

  t.Log(rec.Body.String())

  if rec.Code != http.StatusOK {
  	t.Error("Got ", rec.Code)
  }

  if !strings.Contains(rec.Body.String(), "Hello Yose") {
    t.Error("Does not contain expected string")
  }

}