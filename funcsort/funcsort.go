package main

import "sort"
import "fmt"

// To sort by a custom function in Go, there needs to be a corresponding type.
// The ByLength type was created for this purpose.
type ByLength []string

// Implement th esort.Interface: Len, less, and Swap on the ByLength type so
// we can use the sort package's generic Sort function.
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"apple", "peach", "mango", "pear", "chaimet", "pommecythere"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

