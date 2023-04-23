package templates

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates(pattern string) {
	templates = template.Must(template.ParseGlob(pattern))
}

func ExecuteTemplates(writer http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(writer, template, data)
}
