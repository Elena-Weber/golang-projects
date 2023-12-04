package main

import (
  "math/rand" // random number generator
	// if randomization doesn't work properly, uncomment line below
	// "time"
	"log"
)

const numPool = 100

func calculateValue(intChannel chan int){
	// if randomization doesn't work properly, uncomment line below
	// rand.Seed(time.Now().UnixNano())

	randomNumber := randNum(numPool)

	// push random number into channel
	intChannel <- randomNumber
}

// what it takes in and what returns
func randNum(n int) int {
	value := rand.Intn(n)
	return value
}

func main() {
	intChannel := make(chan int)

	// execute closing this function after it's done
	defer close(intChannel)

	go calculateValue(intChannel)

	// save data from channel into var as soon as it arrives
	num := <-intChannel
	log.Println(num)
}