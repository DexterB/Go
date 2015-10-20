package main

import (
	"fmt"
	"net/http"
    "time"

)

// List of resources to get concurrently
var urls = []string{
	"http://www.ruby.com/",
	"http://golang.org/",
	"http://matt.aimonetti.net/",
}

// HttpResponse type encapsulates a response.
type HttpResponse struct {
	url string
	response *http.Response
    err error
}

// asyncHttpGets gets a set of URLs in parallel and returns the responses.
func asyncHttpGets(urlsIn []string) []*HttpResponse {
	// Create an instance of the channel.
	ch := make(chan *HttpResponse)

	// Create an empty instance of a slice containing pointers to
	// HttpResponse structures.
	responses := []*HttpResponse{}

	// Iterate through the URLs and fetch the resources associated with them.
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			resp.Body.Close()
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}

	// Poll for channel events.
	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		case <- time.After(50 * time.Millisecond):
            // Tick every 50 millisecond interval.
			fmt.Printf(".")
		}
	}

	// Return the responses.
	return responses
}

func main() {
	results := asyncHttpGets(urls)
	for _, result := range results {
		fmt.Printf("%s status: %s\n", result.url, result.response.Status)
	}
}