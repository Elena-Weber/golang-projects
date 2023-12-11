package main

import (
	"fmt"
	"go-web-app/pkg/config"
	"go-web-app/pkg/handlers"
	"go-web-app/pkg/render"
	"log"
	"net/http"
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

	// Sprtintf = string print with variable type string
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))

	// create server with port and routes
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	// listen for errors on port http://localhost:8080/
	err = srv.ListenAndServe()
	log.Fatal(err)
}