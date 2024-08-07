// logify/logger.go
package logify

import (
	"fmt"
	"time"
)

// Logger is a structure for managing logging.
type Logger struct {
	module   string
	consumer string
	env      string // Environment, e.g., "production" or "development"
}

// NewLogger creates a new instance of Logger with the specified module, consumer, and environment.
func NewLogger(module, consumer, env string) *Logger {
	return &Logger{
		module:   module,
		consumer: consumer,
		env:      env,
	}
}

// Log prints a formatted message to standard output.
func (l *Logger) Log(logType LogType, message string) {
	color, logTypeString := LogColor[logType], logType.logTypeParser()
	fmt.Printf("[%s%s%s] [%s] [%s:%s] %s\n",
		color,
		logTypeString,
		Reset,
		time.Now().Format(time.RFC3339),
		l.consumer,
		l.env,
		message,
	)
}

// logTypeParser converts LogType to a readable string.
func (lt LogType) logTypeParser() string {
	switch lt {
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case DEBUG:
		return "DEBUG"
	default:
		return "UNKNOWN"
	}
}
