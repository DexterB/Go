//Pipeline1 demonstrates an infinite 3-stage pipeline
package main

import "fmt"

//+
func main() {
	naturals := make(chan int)
	squares  := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <- naturals
			squares <- x * x
		}
	}()

	// Printer (in main go routine)
	for {
		fmt.Println(<-squares)
	}
}
//-
