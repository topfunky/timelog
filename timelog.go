package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gosuri/uiprogress"
)

func main() {
	boolPtr := flag.Bool("report", false, "a bool")
	flag.Parse()

	filename := flag.Arg(0)

	if *boolPtr {
		week := Week(filename)
		fmt.Println(RenderReportForWeek(week))
	} else {
		doProgress(filename, (45*60)/15, 15)
	}
}

func doProgress(filename string, cycles int, secondsPerCycle time.Duration) {
	uiprogress.Start()               // start rendering
	bar := uiprogress.AddBar(cycles) // Add a new bar

	bar.Fill = '+'

	// optionally, append and prepend completion and elapsed time
	bar.AppendElapsed()

	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return "Task"
	})

	bar.AppendFunc(func(b *uiprogress.Bar) string {
		week := Week(filename)
		return "\n" + RenderReportForWeek(week)
	})

	for bar.Incr() {
		time.Sleep(time.Second * secondsPerCycle)
	}
}
