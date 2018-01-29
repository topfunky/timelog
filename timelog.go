package main

import (
	"time"

	"github.com/gosuri/uiprogress"
)

func main() {
	doProgress((45*60)/15, 15)
}

func doProgress(cycles int, secondsPerCycle time.Duration) {
	uiprogress.Start()               // start rendering
	bar := uiprogress.AddBar(cycles) // Add a new bar

	bar.Fill = '+'

	// optionally, append and prepend completion and elapsed time
	bar.AppendElapsed()

	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return "Task"
	})

	bar.AppendFunc(func(b *uiprogress.Bar) string {
		return "\n\nREPORT GOES HERE " + time.Now().String()
	})

	for bar.Incr() {
		time.Sleep(time.Second * secondsPerCycle)
	}
}
