package main

import "time"
import "fmt"

func main() {
	// Timers represetn a single event in the future.  Tell the time how long
	// you want it to wait, and it provides a channel that will be notified
	// at that time.

	timer1 := time.NewTimer(2 * time.Second)

	// <- blocks on the timer's channel C until it sends a value indicating
	// that the timer has expired.
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// To just wait use timeSleep. One reason a timer may be useful is that
	// it can be cancelled before it expires.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
