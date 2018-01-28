package main

import (
	"testing"
)

func TestTaskListFromFilename(t *testing.T) {
	taskList := TaskList("test/fixtures/todo.md")
	if len(taskList.tasks) != 11 {
		t.Error("Expected: 11, Got:", len(taskList.tasks))
	}
}
