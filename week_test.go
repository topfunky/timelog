package main

import (
	"testing"
)

func TestWeekFirstDay(t *testing.T) {
	week := Week("test/fixtures/todo.md")
	monday := week.tasklists[0]
	if len(monday.tasks) != 9 {
		t.Error("Expected: 9, Got:", len(monday.tasks))
	}
	if monday.title != "Monday" {
		t.Error("Expected: 'Monday', Got:", monday.title)
	}
}

func TestWeekSecondDay(t *testing.T) {
	week := Week("test/fixtures/todo.md")
	tuesday := week.tasklists[1]
	if len(tuesday.tasks) != 5 {
		t.Error("Expected: 5, Got:", len(tuesday.tasks))
	}
	if tuesday.title != "Tuesday" {
		t.Error("Expected: 'Tuesday', Got:", tuesday.title)
	}
}
