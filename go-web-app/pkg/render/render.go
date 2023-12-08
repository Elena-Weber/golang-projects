package render

import (
	"bytes"
	"go-web-app/pkg/config"
	"go-web-app/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// this var is a pointer to config.AppConfig
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
  app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
  return td
}

// ResponseWriter writes to browser
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// get template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache() // ignore error by using _
	}
  

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get a template from template cache")
	}

	// fine tuning error search
	buf := new(bytes.Buffer)

  td = AddDefaultData(td)

	_ = t.Execute(buf, td)

  _, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// these 2 lines are identical and create a map of strings of template pointers
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

  // get all template files as a slice of strings
	pages, err := filepath.Glob("./templates/*.page.tmpl")
  if err != nil {
		return myCache, err
	}

  // map through collected templates, _ means we don't care about index
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// get all layout files as a slice of strings
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	
	return myCache, nil
}