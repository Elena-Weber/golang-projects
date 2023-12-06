package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) { // pointer to request
	f, err := divideValues(100.0, 5.0)
	if err != nil {
		fmt.Fprint(w, "Cannot divide by 0")
		return
	}

	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 5.0, f)) // string print with variable type float32
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}

// main is the main app function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	// Sprtintf = string print with variable type string
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))

	//ignore errors, listen to port http://localhost:8080/
	_ = http.ListenAndServe(portNumber, nil)
}