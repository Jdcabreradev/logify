package logify

//TODO: fix README.md for the v2.0.0 release
import (
	"fmt"
	"os"
	"time"
)

// Logger is a structure for managing logging.
// The Logger structure contains all the information and functionality needed to handle logging.
type Logger struct {
	module      string   // Name of the module generating the logs.
	consumer    string   // Name of the consumer or part of the application using the logger.
	env         string   // Environment in which the application is running (e.g., production, development).
	logDirPath  string   // Directory path where log files will be saved.
	logFilePath string   // Full path of the current log file.
	logFile     *os.File // Pointer to the open log file.
}

// NewLogger creates a new instance of Logger with the specified module, consumer, environment, and log directory.
// Creates a new instance of Logger with the specified module, consumer, environment, and log directory.
func NewLogger(module, consumer, env, logDirPath string) *Logger {
	// Ensure the log directory exists
	// Make sure the log directory exists. If not, create it.
	if err := os.MkdirAll(logDirPath, 0755); err != nil {
		// Print an error message if the log directory cannot be created.
		fmt.Printf("Error creating log directory: %v\n", err)
		return nil
	}

	// Format a timestamp in the desired format for the log file name.
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	// Build the full log file path using the directory and timestamp.
	logFilePath := fmt.Sprintf("%s/app_%s.log", logDirPath, timestamp)

	// Open the log file for writing, creating it if it does not exist, and appending to it if it does.
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// Print an error message if the log file cannot be opened.
		fmt.Printf("Error opening log file: %v\n", err)
		return nil
	}

	// Return a new Logger instance with the specified parameters and the open log file.
	return &Logger{
		module:      module,
		consumer:    consumer,
		env:         env,
		logDirPath:  logDirPath,
		logFilePath: logFilePath,
		logFile:     file,
	}
}

// ConditionalLog checks a condition and calls the Log method if the condition is true.
// This wrapper function helps in deciding whether to log a message based on a conditional check.
func (l *Logger) ConditionalLog(shouldLog bool, logType LogType, message string) {
	// Check the condition.
	if shouldLog {
		// If the condition is true, call the original Log method.
		l.Log(logType, message)
	}
}

// Log prints a formatted message to standard output and saves to file if saveLogs is configured.
// Prints a formatted message to the standard output and saves it to the file if configured.
func (l *Logger) Log(logType LogType, message string) {
	// Get the color and string representation for the current log type.
	color, logTypeString := LogColor[logType], logType.logTypeParser()

	// Create a log message with colors for console output.
	logMessage := fmt.Sprintf("[%s%s%s] [%s] [%s:%s] %s\n",
		color,
		logTypeString,
		Reset,
		time.Now().Format(time.RFC3339),
		l.consumer,
		l.env,
		message,
	)

	// Create a log message without colors for file output.
	logMessageNoColor := fmt.Sprintf("[%s] [%s] [%s:%s] %s\n",
		logTypeString,
		time.Now().Format(time.RFC3339),
		l.consumer,
		l.env,
		message,
	)

	// Print the colored message to the standard output (console).
	fmt.Print(logMessage)

	// If the log file path is configured and the file is open, write the message without colors to the file.
	if l.logFilePath != "" && l.logFile != nil {
		_, err := l.logFile.WriteString(logMessageNoColor)
		if err != nil {
			// Print an error message if the log file cannot be written to.
			fmt.Printf("Error writing to log file: %v\n", err)
		}
	}
}

// Close closes the log file if it is open.
// Closes the log file if it is open.
func (l *Logger) Close() {
	if l.logFile != nil {
		// Attempt to close the log file.
		err := l.logFile.Close()
		if err != nil {
			// Print an error message if the log file cannot be closed.
			fmt.Printf("Error closing log file: %v\n", err)
		}
	}
}

// logTypeParser converts LogType to a readable string.
// Converts LogType to a readable string representation.
func (lt LogType) logTypeParser() string {
	switch lt {
	case INFO:
		return "INFO" // String representation for INFO log type.
	case WARNING:
		return "WARNING" // String representation for WARNING log type.
	case ERROR:
		return "ERROR" // String representation for ERROR log type.
	case DEBUG:
		return "DEBUG" // String representation for DEBUG log type.
	default:
		return "UNKNOWN" // String representation for unknown log types.
	}
}
