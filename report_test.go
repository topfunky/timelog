package main

import (
	"fmt"
	"testing"
)

func TestReportTaskList(t *testing.T) {
	week := Week("test/fixtures/todo.md")
	monday := week.tasklists[0]
	text := RenderReportForTaskList(monday)
	if text == "" {
		t.Error("Expected report to have content but was: ", text)
	}
}

func TestReportWeek(t *testing.T) {
	week := Week("test/fixtures/todo.md")
	text := RenderReportForWeek(week)
	if text == "" {
		t.Error("Expected report to have content but was: ", text)
	}
}

func TestReportPartialWeek(t *testing.T) {
	week := Week("test/fixtures/partial.md")
	text := RenderReportForWeek(week)
	if text == "" {
		t.Error("Expected report to have content but was: ", text)
	}
	fmt.Println(text)
}
