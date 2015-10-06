package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)


// In previous examples we used explicit looking with mutexes to synchronize
// access to shared state across multiple goroutines. Another option is to
// use the build-in synchronization feature of goroutines and channels to
// achieve the same result. This channel-based approach aligns with Go's
// ideas of sharing memory by communicating and having each piece of data
// owned by exactly 1 goroutine.


// In this example, the state is owned by a single goroutine. This will
// guarantee that the data is never corrupted with concurrent access.  To
// read or write state, other goroutine will send messages to the owning go
// routine and receive corresponding replies. These readOp and writeOp
// structures encapsulate those requests and a way for the owning goroutine
// to respond.

type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {

	var ops int64 = 0 // Count how many operations we perform.

	// The read and write channels to be used by other goroutines.
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// This is the goroutine which owns the state, which is a map as in the
	// previous example but now private to the stateful goroutine.
	go func() {
		state := make(map[int] int)
		for {
			select {
			case read := <-reads: read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp {
							key: rand.Intn(5),
							resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
}
