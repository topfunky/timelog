package main

import "strings"

// TL stores a list of task structs.
type TL struct {
	title    string
	tasks    []T
	duration int
}

// TaskList manages a file of task items and times.
func TaskList(rawTitle string, lines []string) (tl TL) {
	// Parse title from raw title "# Monday"
	tl.title = strings.Join(strings.Split(rawTitle, " ")[1:], " ")
	// Append raw text line to tasks and calculate duration
	// based on timestamp of previous task
	for index, line := range lines {
		tl.tasks = append(tl.tasks, Task(line))
		if index > 0 {
			tl.tasks[index].duration = tl.tasks[index].timestamp - tl.tasks[index-1].timestamp
			tl.duration += tl.tasks[index].duration
		}
	}
	return
}
