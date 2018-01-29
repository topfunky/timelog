package main

import (
	"testing"
)

func makeTaskList() TL {
	lines := []string{
		"- [x] BOD 08:00",
		"- [ ] Make bacon #cooking 09:00",
		"- [ ] Reply to emails #communication 09:30",
		"- [ ] Call Krusty the Clown #communication 12:00",
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
	if taskList.duration != 240 {
		t.Error("Expected duration to be 240 but it was", taskList.duration)
	}
}

func TestTaskListCalculatesTaskDuration(t *testing.T) {
	taskList := makeTaskList()
	if taskList.tagStats["#communication"] != 180 {
		t.Error("Expected tag duration to be 180 but it was", taskList.tagStats["#communication"])
	}
}
