package logging

import (
	"fmt"
	"time"

	"github.com/logrusorgru/aurora"
)

const (
	CRITICAL = 5
	ERROR    = 4
	WARNING  = 3
	INFO     = 2
	DEBUG    = 1
)

type Logger struct {
	Level int
	name  string
}

func NewLogger(name string) Logger {
	return Logger{DEBUG, name}
}

func (l *Logger) SetLevel(level int) {
	l.Level = level
}

func (l *Logger) log(level int, message string) {
	if level < l.Level {
		return
	} else {
		var level_string string
		var colored_template aurora.Value
		log_template := "%s - %s %s %s"
		switch level {
		case CRITICAL:
			level_string = "CRITICAL"
			colored_template = aurora.BrightRed(log_template)
		case ERROR:
			level_string = "ERROR"
			colored_template = aurora.BrightRed(log_template)
		case WARNING:
			level_string = "WARNING"
			colored_template = aurora.BrightYellow(log_template)
		case INFO:
			level_string = "INFO"
			colored_template = aurora.BrightCyan(log_template)
		case DEBUG:
			level_string = "DEBUG"
			colored_template = aurora.BrightGreen(log_template)
		default:
			level_string = "NOTSET"
			colored_template = aurora.BrightWhite(log_template)
		}
		ts := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(aurora.Sprintf(
			colored_template,
			l.name,
			ts,
			level_string,
			message,
		))
	}
}

func (l *Logger) Critical(message string) {
	l.log(CRITICAL, message)
}

func (l *Logger) Error(message string) {
	l.log(ERROR, message)
}

func (l *Logger) Warning(message string) {
	l.log(WARNING, message)
}

func (l *Logger) Info(message string) {
	l.log(INFO, message)
}

func (l *Logger) Debug(message string) {
	l.log(DEBUG, message)
}
