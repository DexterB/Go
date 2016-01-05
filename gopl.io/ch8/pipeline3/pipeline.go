// Pipeline3 demonstrates an infinite 3-stage pipeline with range, close, and
// unidirectional channel types.
package main

import "fmt"

//!+
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for w := range in {
		fmt.Println(w)
	}
}

func main() {
	naturals := make(chan int)
	squares  := make(chan int)

	// Counter
	go counter (naturals)
	go squarer (squares, naturals)
	printer(squares)
}

//!-
