package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

// capital letter function name means it's accessible from outside
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

 // About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	// parse templates
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	// handle errors from parsed template
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing a template:", err)
		return
	}
}

// main is the main app function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// Sprtintf = string print with variable type string
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))

	//ignore errors, listen to port http://localhost:8080/
	_ = http.ListenAndServe(portNumber, nil)
}