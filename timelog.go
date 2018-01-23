package main

import (
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	time.Sleep(10 * time.Second)                                 // Run for some time to simulate work
	s.Stop()
}
