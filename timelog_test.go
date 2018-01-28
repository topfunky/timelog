package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type T struct {
	raw         string
	minutes     int
	description string
}

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

func TestNewUncompleted(t *testing.T) {
	task := Task("- [ ] Make bacon 09:34")
	if task.description != "Make bacon" {
		t.Error("Expected: Make bacon, Got:", task.description)
	}
	if task.minutes != 574 {
		t.Error("Expected: 574, Got:", task.minutes)
	}
}

func TestNewCompleted(t *testing.T) {
	task := Task("- [x] Make bacon 09:35")
	if task.description != "Make bacon" {
		t.Error("Expected: Make bacon", task.description)
	}
	if task.minutes != 575 {
		t.Error("Expected: 575, Got:", task.minutes)
	}
}

// https://github.com/fatih/color

// Parse file: -(.*)\s(\d\d:\d\d)
// Convert time to seconds
// Pretty print seconds to HH:MM
// Sort lines by time
// Calculate time since previous timed event
// Print report of time spent

// Parse tag #course
// Print report by time on tag
// Use terminal library to print running time and report
