package tests

import (
	"testing"

	"github.com/jpartogi/goyose/config"
)

var configuration = &config.Configuration{}

func TestParseProdConfig(t *testing.T) {
  config.Load("../config/config.json", configuration)

  templateFolder:=configuration.View.TemplateFolder

  t.Log("TemplateFolder: "+ templateFolder)

  if templateFolder != "templates" {
    t.Error("Does not contain expected string")
  }
}