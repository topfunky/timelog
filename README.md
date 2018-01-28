# Time Log

Reads a Markdown task file with timestamps and generates a report.

## Planned Features

- Show a spinner while running
- Show HH:MM count up
- Render spinner in green for first X minutes, yellow for a few, then red if over a threshold
- Show a running report by tag

### Notes

// https://github.com/fatih/color

// Parse file: -(.*)\s(\d\d:\d\d)
// Convert time to seconds
// Pretty print seconds to HH:MM
// Sort lines by time
// Calculate time since previous timed event
// Print report of time spent

// Parse tag #course
// Print report by time on tag
// Use terminal library to print running time and report
