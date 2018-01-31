package main

import "strings"

// TL stores a list of task structs.
type TL struct {
	title    string
	tasks    []T
	duration int
	tagStats map[string]int
}

// TaskList manages a file of task items and times.
func TaskList(rawTitle string, lines []string) (tl TL) {
	// Parse title from raw title "# Monday"
	tl.title = strings.Join(strings.Split(rawTitle, " ")[1:], " ")
	tl.tagStats = make(map[string]int)
	// Append raw text line to tasks and calculate duration
	// based on timestamp of previous task
	for index, line := range lines {
		task := Task(line)
		tl.tasks = append(tl.tasks, task)
		if index > 0 {
			if task.timestamp == 0 {
				continue
			}
			tl.tasks[index].duration = tl.tasks[index].timestamp - tl.tasks[index-1].timestamp
			tl.duration += tl.tasks[index].duration
			// TODO: Calculate stats for all tags
			if len(tl.tasks[index].tags) > 0 {
				// Explicit tag
				tl.tagStats[tl.tasks[index].tags[0]] += tl.tasks[index].duration
			} else if tl.tasks[index].duration > 0 {
				tl.tagStats["#other"] += tl.tasks[index].duration
			}
		}
	}
	return
}
