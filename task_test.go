package main

import (
	"testing"
)

func TestTaskUncompleted(t *testing.T) {
	task := Task("- [ ] Make bacon 09:34")
	if task.description != "Make bacon" {
		t.Error("Expected: Make bacon, Got:", task.description)
	}
	if task.timestamp != 574 {
		t.Error("Expected: 574, Got:", task.timestamp)
	}
}

func TestTaskCompleted(t *testing.T) {
	task := Task("- [x] Make bacon 09:35")
	if task.description != "Make bacon" {
		t.Error("Expected: Make bacon", task.description)
	}
	if task.timestamp != 575 {
		t.Error("Expected: 575, Got:", task.timestamp)
	}
}

func TestTaskCompletedWithTimeError(t *testing.T) {
	task := Task("- [x] Make bacon")
	if task.timestamp != 0 {
		t.Error("Expected: 0, Got:", task.timestamp)
	}
}
