package main

import "fmt"

// This example uses for and range to iterate through values received from a
// channel.
//
func main() {

	queue := make(chan string, 2);
	queue <- "one"
	queue <- "two"
	close(queue);

	for elem := range queue {
		fmt.Println(elem);
	}

	err := true;
	var one, two string
	one, err = <- queue
	two, err = <- queue
	fmt.Println(err)

	if !err {
		fmt.Print("Nothing in queue")
	} else {
		fmt.Println("First value: ", one, "Second value: ", two)
	}
}
