package main

import (
	"io/ioutil"
	"log"
	"strings"
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
