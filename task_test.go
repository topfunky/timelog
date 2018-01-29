package main

import (
	"testing"
)

func makeUncompletedTask() T {
	return Task("- [ ] Make bacon #eat 09:34")
}

func makeCompletedTask() T {
	return Task("- [x] Make bacon #cooking 09:35")
}

func TestTaskUncompleted(t *testing.T) {
	task := makeUncompletedTask()
	if task.description != "Make bacon" {
		t.Error("Expected: Make bacon, Got:", task.description)
	}
	if task.timestamp != 574 {
		t.Error("Expected: 574, Got:", task.timestamp)
	}
}

func TestTaskCompletedHasTitle(t *testing.T) {
	task := makeCompletedTask()
	if task.description != "Make bacon" {
		t.Error("Expected: 'Make bacon', Got:", task.description)
	}
}

func TestTaskCompletedHasTimestamp(t *testing.T) {
	task := makeCompletedTask()
	if task.timestamp != 575 {
		t.Error("Expected: 575, Got:", task.timestamp)
	}
}

func TestTaskCompletedHasTags(t *testing.T) {
	task := makeCompletedTask()
	if len(task.tags) < 1 || task.tags[0] != "#cooking" {
		t.Error("Expected tag: ['#cooking'], Got:", task.tags)
	}
}

func TestTaskCompletedWithTimeError(t *testing.T) {
	task := Task("- [x] Make bacon")
	if task.timestamp != 0 {
		t.Error("Expected: 0, Got:", task.timestamp)
	}
}
