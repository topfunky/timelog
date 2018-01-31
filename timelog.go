package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gosuri/uiprogress"
)

func main() {
	boolPtr := flag.Bool("report", false, "a bool")
	versionPtr := flag.Bool("version", false, "Display the version number of the app")
	flag.Usage = func() {
		printVersionInformation()
		fmt.Fprintln(os.Stderr, "usage:")
		flag.PrintDefaults()
	}
	flag.Parse()

	filename := flag.Arg(0)

	if *boolPtr {
		week := Week(filename)
		fmt.Println(RenderReportForWeek(week))
	} else if *versionPtr {
		printVersionInformation()
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

func printVersionInformation() {
	fmt.Fprintf(os.Stderr, "%s version %s\n", os.Args[0], VERSION)
	fmt.Fprintf(os.Stderr, "built %s\n", BUILD_DATE)
}
