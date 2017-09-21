package tests

import (
	"net/http"
	"net/http/httptest"

 	"github.com/jpartogi/goyose/view"
)

func TestHtmlPage(url string, method string) (*http.Request, *httptest.ResponseRecorder) {
 	config:=view.Config{ TemplateFolder: "../templates" }
  	view.Configure(config)

	req := httptest.NewRequest(method, url, nil)
  	req.Header.Set("Content-Type", "text/html")
  	rec := httptest.NewRecorder()

  	return req, rec
}

func TestAPI(url string, method string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	return req, rec
}