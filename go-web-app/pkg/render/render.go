package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// ResponseWriter writes to browser
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	// fine tuning error search
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
  if err != nil {
		log.Println(err)
	}

  _, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
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