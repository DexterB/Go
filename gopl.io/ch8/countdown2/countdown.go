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
	select {
	case <-time.After(10 * time.Second):
		// Do nothing.
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

//!-
