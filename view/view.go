package view

import (
	"html/template"
	"net/http"
	"log"
	"os"
)

var (
	view View
	config Config
)

type Config struct {
	TemplateFolder string `json:"TemplateFolder"`	
}

// View attributes
type View struct {
	Template string
}

// Configure sets the view information
func Configure(c Config) {
	config = c
}

// New returns a new view
func New(template string) *View {
	v := &View{}

	v.Template = template

	return v
}

// Render renders a template to the writer
func (v *View) Render(w http.ResponseWriter) {
	log.Println(config.TemplateFolder)

	t, err := template.ParseFiles(config.TemplateFolder+string(os.PathSeparator)+v.Template)

	if err != nil {
		http.Error(w, "Template File Error: "+err.Error(), http.StatusInternalServerError)
	}

	t.Execute(w, nil)
}

