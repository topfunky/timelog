package main

import "strings"

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
