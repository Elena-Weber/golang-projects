package handlers

import (
	"go-web-app/pkg/config"
	"go-web-app/pkg/render"
	"go-web-app/pkg/models"
	"net/http"
)

// global var of type Repository
// Repo repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
 // NewHandlers sets repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is home page handler
// Capital letter function name means it's global
// receives a receiver of type Repository
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{}) // send TemplateData to template
}

 // About is about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
  stringMap := make(map[string]string)
	stringMap["test"] = "Hello there!"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{ // send TemplateData to template
		StringMap: stringMap,
	})
}