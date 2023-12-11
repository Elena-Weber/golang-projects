package main

import (
	"go-web-app/pkg/config"
	"net/http"
	// "github.com/bmizerany/pat"
	"go-web-app/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// default server mux
  mux := chi.NewRouter()

	// add chi middleware to handle app panics
	mux.Use(middleware.Recoverer)

	// use middlewares we created in middleware.go
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)

	// routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux

	// THIS IS PAT APPROACH
	// // default server mux
  // mux := pat.New()
	// // routes
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	// return mux
}