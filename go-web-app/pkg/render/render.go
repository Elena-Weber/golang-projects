package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// parse templates
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	// handle errors from parsed template
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing a template:", err)
		return
	}
}