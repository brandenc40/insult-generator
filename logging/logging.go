package logging

import (
	"fmt"
	"time"

	"github.com/logrusorgru/aurora"
)

// Log Levels
const (
	DEBUG = iota + 1
	INFO
	WARNING
	ERROR
	CRITICAL
)

// Logger -
type Logger struct {
	Level int
	Name  string
}

// NewLogger -
func NewLogger(name string) Logger {
	return Logger{DEBUG, name}
}

// SetLevel -
func (l *Logger) SetLevel(level int) {
	l.Level = level
}

func (l *Logger) log(level int, message string) {
	if level < l.Level {
		return
	}
	var levelString string
	var coloredTemplate aurora.Value
	logTemplate := "%s - %s %s %s"
	switch level {
	case CRITICAL:
		levelString = "CRITICAL"
		coloredTemplate = aurora.BrightRed(logTemplate)
	case ERROR:
		levelString = "ERROR"
		coloredTemplate = aurora.BrightRed(logTemplate)
	case WARNING:
		levelString = "WARNING"
		coloredTemplate = aurora.BrightYellow(logTemplate)
	case INFO:
		levelString = "INFO"
		coloredTemplate = aurora.BrightCyan(logTemplate)
	case DEBUG:
		levelString = "DEBUG"
		coloredTemplate = aurora.BrightGreen(logTemplate)
	default:
		levelString = "NOTSET"
		coloredTemplate = aurora.BrightWhite(logTemplate)
	}
	ts := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(aurora.Sprintf(
		coloredTemplate,
		l.Name,
		ts,
		levelString,
		message,
	))
}

// Critical -
func (l *Logger) Critical(message string) {
	l.log(CRITICAL, message)
}

// Error -
func (l *Logger) Error(message string) {
	l.log(ERROR, message)
}

// Warning -
func (l *Logger) Warning(message string) {
	l.log(WARNING, message)
}

// Info -
func (l *Logger) Info(message string) {
	l.log(INFO, message)
}

// Debug -
func (l *Logger) Debug(message string) {
	l.log(DEBUG, message)
}
