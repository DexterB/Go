// Countdown implements the countdown for a rocket launch. The launch can be
// aborted at anytime by pressing any key.
package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Lift off!")
}

//!+
func main() {
	// Create an abort channel.

	//!-

	//!+abort
	abort := make(chan struct{})
	go func () {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	//!-abort

	//!+
	fmt.Println("Commencing countdown. Press return to abort.")
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-ticker.C:
			fmt.Println(countdown)
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

//!-

