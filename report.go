package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
)

// RenderReportForWeek returns a textual report for a full week.
func RenderReportForWeek(week W) (report string) {
	for _, tasklist := range week.tasklists {
		report += RenderReportForTaskList(tasklist)
	}
	return
}

// RenderReportForTaskList returns a textual report
// from a task list's tags and durations.
func RenderReportForTaskList(taskList TL) string {

	// Monday:
	//   #communication   :35
	//   #meetings       2:35
	//   total           3:05
	b := new(bytes.Buffer)
	w := tabwriter.NewWriter(b, 16, 0, 2, ' ', tabwriter.FilterHTML)

	for tag, duration := range taskList.tagStats {
		fmt.Fprintln(w, strings.Join([]string{
			"  " + tag,
			formatMinutes(duration)}, "\t")+"\t")
	}
	w.Flush()
	return strings.Join([]string{taskList.title, b.String()}, "\n")
}

func formatMinutes(minutes int) string {
	hours := minutes / 60
	remainder := minutes % 60
	return fmt.Sprintf("%d:%02d", hours, remainder)
}
