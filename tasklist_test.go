package main

import (
	"fmt"
	"testing"
)

func makeTaskList() TL {
	lines := []string{
		"- [x] BOD 08:00",
		"- [ ] Make bacon 09:00",
		"- [ ] Reply to emails 09:30",
		"- [ ] Call Krusty the Clown 12:00",
	}
	return TaskList("# Monday", lines)
}

func TestTaskListMakesTasks(t *testing.T) {
	taskList := makeTaskList()
	if len(taskList.tasks) != 4 {
		t.Error("Expected: 4, Got:", len(taskList.tasks))
	}
}

func TestTaskListCalculatesDurationOfFirstTask(t *testing.T) {
	taskList := makeTaskList()
	if taskList.tasks[0].duration != 0 {
		t.Error("Expected BOD task to have duration 0 but it was", taskList.tasks[0].duration)
	}
}

func TestTaskListCalculatesDurationOfOtherTask(t *testing.T) {
	taskList := makeTaskList()
	if taskList.tasks[1].duration != 60 {
		t.Error("Expected 'Make bacon' task to have duration 60 but it was", taskList.tasks[1].duration)
	}
}

func TestTaskListCalculatesTotalDuration(t *testing.T) {
	taskList := makeTaskList()
	fmt.Println(taskList)
	if taskList.duration != 240 {
		t.Error("Expected duration to be 240 but it was", taskList.duration)
	}
}
