package main

import (
	"fmt"
	"testing"
)

func TestReportTaskList(t *testing.T) {
	week := Week("test/fixtures/todo.md")
	monday := week.tasklists[0]
	RenderReportForTaskList(monday)
}

func TestReportWeek(t *testing.T) {
	week := Week("test/fixtures/todo.md")
	text := RenderReportForWeek(week)
	fmt.Println(text)
}
