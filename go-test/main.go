package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100.0, 10.0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Result is", result)
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 { // create own error condition and message
		return result, errors.New("Cannot divide by zero")
	}

	result = x / y
	return result, nil
}