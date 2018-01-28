package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gosuri/uiprogress"
)

// W stores data for a single week of tasks
type W struct {
	tasklists []TL
	// TODO: filename string
}

// Week manages a file of task items and times.
func Week(filename string) (week W) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(bytes), "\n")

	rawLines := make([]string, 0)
	rawTitle := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			if len(rawLines) > 0 {
				week.tasklists = append(week.tasklists, TaskList(rawTitle, rawLines))
			}
			rawTitle = line
			// TODO: Use existing section or reset
			// week.tasklists = append(week.tasklists, TaskList(rawTitle, lines))
			// week.tasklists[0].tasks = append(week.tasklists[0].tasks, Task(line))

		} else if strings.HasPrefix(line, "- [") {
			rawLines = append(rawLines, line)
		}
	}
	return
}

// TL stores a list of task structs.
type TL struct {
	title string
	tasks []T
}

// TaskList manages a file of task items and times.
func TaskList(rawTitle string, lines []string) (tl TL) {
	// Parse title from raw title "# Monday"
	tl.title = strings.Join(strings.Split(rawTitle, " ")[1:], " ")
	for _, line := range lines {
		tl.tasks = append(tl.tasks, Task(line))
	}
	return
}

// T is the data structure for a single task
type T struct {
	minutes     int
	description string
}

// Task creates a new task consisting of a description and a number of minutes
func Task(raw string) (t T) {
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
	doProgress((45*60)/15, 15)
	// doSpinner()
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

	for bar.Incr() {
		time.Sleep(time.Second * secondsPerCycle)
	}
}
