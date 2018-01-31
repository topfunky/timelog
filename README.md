# Time Log

A personal project to read a Markdown task file with timestamps and generate a report.

Also displays a running 45 minute progress bar.

## Usage

Call timelog with the name of a Markdown task file.

    timelog ~/taskfile.md

    Task [+++>------------]    5m15s
    Monday 3:42
      #break         1:14
      #code          1:07
      #blog          0:35

## File format

The task log should be a Markdown file with GitHub style checkboxes and an HH:MM timestamp at the end of the line.

Hashtags can optionally be used to categorize tasks. (Only one hashtag per line is supported.)

    # Monday
    
    - [x] Start day 08:45
    - [x] Write code #code 09:00
    - [ ] Draft blog post
