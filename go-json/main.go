package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Married bool `json:"married"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "John",
			"last_name": "Doe",
			"married": true
		},
		{
			"first_name": "Jane",
			"last_name": "Dope",
			"married": false
		}
	]`
	
	// after Go's functions: marshal and unmarshal
	var unmarshalled []Person // slice

	// turns json into struct of values, removes keys
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Kelly"
	m1.LastName = "Peterson"
	m1.Married = false

	var m2 Person
	m2.FirstName = "Michael"
	m2.LastName = "Smith"
	m2.Married = true

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	// turns structs into json
	// MarshalIndent is more readable than Marshal
	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("Error marshalling", err)
	}

	fmt.Printf(string(newJson))
}