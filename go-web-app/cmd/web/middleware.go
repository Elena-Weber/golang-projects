package main

import (
	"fmt"
	"net/http"
	"github.com/justinas/nosurf"
)

// WriteToConsole writes to console with every page load and reload
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("User hit the webpage")
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds csrf protection to all post requests
func NoSurf(next http.Handler) http.Handler {
  csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}