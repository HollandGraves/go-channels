package main

import (
	"fmt"
	"net/http"
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

	iterateLink(links, c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!")
		c <- "Might be down I think!"

		go checkLink(link, c)
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep it's up!"
	go checkLink(link, c)
}

func iterateLink(links []string, c chan string) {
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
		iterateLink(links, c)
	}
}
