package main

import (
	"bytes"
	"fmt"
	"sort"
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

	// TODO: Sort by value (duration) with longest at top
	keys := make([]string, 0, len(taskList.tagStats))
	for k := range taskList.tagStats {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, tag := range keys {
		fmt.Fprintln(w, strings.Join([]string{
			"  " + tag,
			formatMinutes(taskList.tagStats[tag])}, "\t")+"\t")
	}
	w.Flush()
	totalDurationFormatted := formatMinutes(taskList.duration)
	return strings.Join([]string{taskList.title + " " + totalDurationFormatted, b.String()}, "\n")
}

func formatMinutes(minutes int) string {
	hours := minutes / 60
	remainder := minutes % 60
	return fmt.Sprintf("%d:%02d", hours, remainder)
}
