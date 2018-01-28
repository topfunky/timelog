package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gosuri/uiprogress"
)

// T is the data structure for a single task
type T struct {
	raw         string
	minutes     int
	description string
}

// Task creates a new task consisting of a description and a number of minutes
func Task(raw string) (t T) {
	t.raw = raw
	t.description = raw
	t.minutes = 0

	t.description = strings.TrimPrefix(t.description, "- [ ] ")
	t.description = strings.TrimPrefix(t.description, "- [x] ")
	if match, _ := regexp.MatchString(" ([0-9][0-9]:[0-9][0-9])$", t.description); match {
		minute, err := strconv.Atoi(t.description[len(t.description)-2:])
		if err != nil {
			fmt.Println(minute)
		}
		hour, err := strconv.Atoi(t.description[len(t.description)-5 : len(t.description)-3])
		if err != nil {
			fmt.Println(err)
		}
		t.minutes = hour*60 + minute
		t.description = t.description[:len(t.description)-6]
	}
	return
}

func main() {
	doProgress()
	// doSpinner()
}

func doProgress() {
	duration := (45 * 60) / 15
	uiprogress.Start()                 // start rendering
	bar := uiprogress.AddBar(duration) // Add a new bar

	// optionally, append and prepend completion and elapsed time
	bar.AppendElapsed()

	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return "Task"
	})

	for bar.Incr() {
		time.Sleep(time.Second * 15)
	}
}
