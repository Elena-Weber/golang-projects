package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w, "Hello World") // prints on webpage
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n)) // prints in console
	})

	//ignore errors, listen to port http://localhost:8080/
	_ = http.ListenAndServe(":8080", nil)
}