package main

import (
	"fmt"
	"io"
	"os"
)

	// to run: go run main.go myfile.txt

func main() {
	// opens file f which is in position 1 in Args slice
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

// prints contents of file f to terminal
	io.Copy(os.Stdout, f)
}