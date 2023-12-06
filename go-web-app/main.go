package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"



// main is the main app function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// Sprtintf = string print with variable type string
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))

	//ignore errors, listen to port http://localhost:8080/
	_ = http.ListenAndServe(portNumber, nil)
}