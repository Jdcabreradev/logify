package logify

import (
	"fmt"
	"time"
)

// LogType is an enumerated type for the different types of logs.
type LogType int

const (
	INFO LogType = iota
	WARNING
	ERROR
	DEBUG
)

// ANSI color codes.
const (
	reset   = "\033[0m"
	black   = "\033[30m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"
)

// logTypeParser converts the log type to a readable string with colors.
func (lt LogType) logTypeParser() (string, string) {
	switch lt {
	case INFO:
		return green, "INFO"
	case WARNING:
		return yellow, "WARNING"
	case ERROR:
		return red, "ERROR"
	case DEBUG:
		return blue, "DEBUG"
	default:
		return reset, "UNKNOWN"
	}
}

// Logger is a structure for managing logging.
type Logger struct {
	module   string
	consumer string
}

// Log prints a color-formatted message to standard output.
func (l *Logger) Log(logType LogType, message string) {
	color, logTypeString := logType.logTypeParser()
	fmt.Printf("[%s%s%s] [%s] [%s] [%s] %s\n",
		color,
		logTypeString,
		reset,
		l.module,
		l.consumer,
		time.Now().Format(time.DateTime),
		message,
	)
}
