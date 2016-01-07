// Countdown implements the countdown for a rocket launch.
package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("Lift off!")
}

//!+
func main() {
	fmt.Println("Commencing countdown")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		<-tick
		fmt.Println(countdown)
	}
	launch()
}

//!-
