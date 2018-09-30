package handlers

import (
	"html/template"
	"net/http"

	"github.com/anabiozz/logger"
)

// RenderTemplate ..
func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		logger.Errorf("Error encountered while parsing the template: %s", err)
	}
	t.Execute(w, templateData)
}
