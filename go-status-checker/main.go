package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"https://google.com",
		"https://facebook.com",
		"https://instagram.com",
		"https://amazon.com",
		"https://golang.org",
	}

	// create channel for data type string
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // go means concurrency, needed for Go Routines
	}

	// iterate through links, each link equals one channel
	// for i := 0; i < len(links); i++ { }

	// for { // same as loop above but infinite (!!!)
	// 	go checkLink(<- c, c)
	// }

	for l := range c { // l equals link from channel
		go func(link string) { // function without name is called 'function literal'
			time.Sleep(10 * time.Second) // timer to check status every 10 seconds
			checkLink(link, c)
		}(l) // passing argument to function literal
	}
}

// it takes 2 args: link as string and c channel for data type string
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Warning!", link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}