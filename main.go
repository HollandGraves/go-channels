package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	c := make(chan string)

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string, c chan string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
