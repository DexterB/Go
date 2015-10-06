package main

import "os"

// A panic means that something went wrong unexpectedly. Mostly it is used to
// fail fast on errors that should not occur during normal operation, aor
// that we are prepared to handle gracefully.

func main() {
	// Panic will be used to check for unexpected errors. This is the only
	// programs in the examples that is designed to panic.
	panic("a problem!")

	// A common use of the panic is to abort if a function returns an error
	// value that we don't know ho wo k or necessarily want to handle.
	_, err := os.Create("/tmp/somebogusfile")
	if (err != nil) {
		panic(err)
	}
}
