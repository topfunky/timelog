package main

import (
	"testing"
)

func TestTaskUncompleted(t *testing.T) {
	task := Task("- [ ] Make bacon 09:34")
	if task.description != "Make bacon" {
		t.Error("Expected: Make bacon, Got:", task.description)
	}
	if task.minutes != 574 {
		t.Error("Expected: 574, Got:", task.minutes)
	}
}

func TestTaskCompleted(t *testing.T) {
	task := Task("- [x] Make bacon 09:35")
	if task.description != "Make bacon" {
		t.Error("Expected: Make bacon", task.description)
	}
	if task.minutes != 575 {
		t.Error("Expected: 575, Got:", task.minutes)
	}
}

func TestTaskCompletedWithTimeError(t *testing.T) {
	task := Task("- [x] Make bacon")
	if task.minutes != 0 {
		t.Error("Expected: 0, Got:", task.minutes)
	}
}

func TestTaskListFromFilename(t *testing.T) {
	taskList := TaskList("test/fixtures/todo.md")
	if len(taskList) != 11 {
		t.Error("Expected: 11, Got:", len(taskList))
	}
}
