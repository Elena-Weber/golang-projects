package main

import (
	"go-web-app/pkg/config"
	"net/http"
	"github.com/bmizerany/pat"
	"go-web-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// default server mux
  mux := pat.New()

	// routes
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}