package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// T is the data structure for a single task
type T struct {
	timestamp   int
	duration    int
	description string
	tags        []string
}

// Task creates a new task consisting of a description and a number of minutes
func Task(raw string) (t T) {
	t.description = raw
	t.timestamp = 0

	t.description = strings.TrimPrefix(t.description, "- [ ] ")
	t.description = strings.TrimPrefix(t.description, "- [x] ")

	// TODO: Try [[:digit:]]{2} or use regexp.Compile
	if match, _ := regexp.MatchString(" ([0-9][0-9]:[0-9][0-9])$", t.description); match {
		minute, err := strconv.Atoi(t.description[len(t.description)-2:])
		if err != nil {
			fmt.Println(minute)
		}
		hour, err := strconv.Atoi(t.description[len(t.description)-5 : len(t.description)-3])
		if err != nil {
			fmt.Println(err)
		}
		t.timestamp = hour*60 + minute
		t.description = t.description[:len(t.description)-6]
	}

	matcher, _ := regexp.Compile("#([a-z]+)")
	if match := matcher.FindString(t.description); match != "" {
		t.tags = append(t.tags, match)
		// Assumes tag is at the end of the description
		t.description = strings.TrimSpace(t.description[:len(t.description)-len(match)])
	}
	return
}
