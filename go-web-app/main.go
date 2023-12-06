package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// capital letter function name means it's accessible from outside
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HOME")
}

 // About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 3)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("ABOUT %d", sum)) // string print with variable type int
}

 // small letter function name means it's private
func addValues(x, y int) int {
	return x + y
}

// main is the main app function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// Sprtintf = string print with variable type string
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))

	//ignore errors, listen to port http://localhost:8080/
	_ = http.ListenAndServe(portNumber, nil)
}