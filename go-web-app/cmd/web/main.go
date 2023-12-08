package main

import (
	"go-web-app/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the main app function
// to run the app - go run cmd/web/*.go - in the root folder of the project
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// Sprtintf = string print with variable type string
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))

	//ignore errors, listen to port http://localhost:8080/
	_ = http.ListenAndServe(portNumber, nil)
}