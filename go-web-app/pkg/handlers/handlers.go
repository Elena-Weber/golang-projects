package handlers

import (
	"go-web-app/pkg/render"
	"net/http"
)

// Capital letter function name means it's accessible from outside
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

 // About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}