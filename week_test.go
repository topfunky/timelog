package main

import (
	"fmt"
	"testing"
)

func TestWeek(t *testing.T) {
	week := Week("test/fixtures/todo.md")
	monday := week.tasklists[0]
	if len(monday.tasks) != 9 {
		t.Error("Expected: 9, Got:", len(monday.tasks))
	}
	if monday.title != "Monday" {
		t.Error("Expected: 'Monday', Got:", monday.title)
	}
	fmt.Println(monday)
}
