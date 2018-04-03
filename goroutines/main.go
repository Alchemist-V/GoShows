package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	wb := []string{
		"http://facebook.com",
		"http://walmart.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	checkSites(wb)

}

func checkSites(wb []string) {
	c := make(chan string)
	for _, site := range wb {
		go checkSiteStatus(site, c)
	}

	// treats channel as something go can iterate upon.
	for l := range c {
		// function literal
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkSiteStatus(link, c)
		}(l)
	}
}

func checkSiteStatus(site string, c chan string) {
	st := time.Now().UnixNano() / int64(time.Millisecond)
	_, err := http.Get(site)
	end := time.Now().UnixNano() / int64(time.Millisecond)
	if err != nil {
		c <- site
		return
	}
	fmt.Println(site, " is working and took "+strconv.FormatInt(end-st, 10)+" millsecs to respond")
	c <- site
}
