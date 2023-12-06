package main

import (
	"net/http"
)

// Capital letter function name means it's accessible from outside
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

 // About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}