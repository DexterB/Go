package main

import "time"
import "fmt"

func main() {

	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on the receive from the limiter channel before serving
	// each request, we limit ourselves to a request every 200 milliseconds.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	// Allow short bursts of requests in the rate limiting scheme while
	// preserving the overall rate limit. This can be accomplished by
	// buffering our limiter channel. this burtyLimiter channel will allow
	// bursts of up to three events.
	for i := 0; i < 3;  i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200 milliseconds, try to add a new value to burstyLimiter, up to
	// its limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 more incoming requests. The first 3 of these will
	// benefit from the burst capability the burstyLimiter.
	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
