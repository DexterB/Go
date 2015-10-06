package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {

	var ops uint64 = 0

	// To simulate concurrent updates, start wiht 50 goroutines that each
	// increment the counter about once a millisecond.
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// Atomically increment th ecounter.
				atomic.AddUint64(&ops, 1)

				// Allow other go routine to proceed.
				runtime.Gosched();
			}
		}()
	}

	// Wait a second to let some operations complete.
	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops);
	fmt.Println("ops", opsFinal)
}
