package main

import (
	"testing"
)

func TestTaskListFromFilename(t *testing.T) {
	lines := []string{
		"- [x] BOD 08:00",
		"- [ ] Make bacon 09:00",
		"- [ ] Reply to emails 09:30",
		"- [ ] Call Krusty the Clown 10:00",
	}
	taskList := TaskList("# Monday", lines)
	if len(taskList.tasks) != 4 {
		t.Error("Expected: 4, Got:", len(taskList.tasks))
	}
}
