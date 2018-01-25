package main

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uiprogress"
)

func main() {
	doProgress()
	// doSpinner()
}

func doSpinner() {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	time.Sleep(10 * time.Second)                                 // Run for some time to simulate work
	s.Stop()
}

func doProgress() {
	duration := (25 * 60) / 15
	uiprogress.Start()                 // start rendering
	bar := uiprogress.AddBar(duration) // Add a new bar

	// optionally, append and prepend completion and elapsed time
	bar.AppendElapsed()

	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return "app: Bacon"
	})

	for bar.Incr() {
		time.Sleep(time.Second * 15)
	}
}
