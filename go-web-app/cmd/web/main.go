package main

import (
	"go-web-app/pkg/render"
	"fmt"
	"go-web-app/pkg/config"
	"go-web-app/pkg/handlers"
	"net/http"
	"log"
)

const portNumber = ":8080"

// main is the main app function
// to run the app - go run cmd/web/*.go - in the root folder of the project
func main() {
	var app config.AppConfig
  
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	// Sprtintf = string print with variable type string
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))

	//ignore errors, listen to port http://localhost:8080/
	_ = http.ListenAndServe(portNumber, nil)
}