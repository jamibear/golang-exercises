package main

import (
	"fmt"
	"net/http"
	"time"
)

// Website status checker
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create channels
	c := make(chan string)

	// spawn go child routines
	for _, link := range links {
		go checkLink(link, c)
	}
	// recieve message from channel
	for l := range c {
		// use function literal for sleep intervals
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// check website status or if website is down
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		// send message through channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// send message through channel
	c <- link
}
