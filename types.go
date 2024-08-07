package logify

// LogType is an enumerated type for the different types of logs.
// LogType is an enumerated type representing the different types of logs.
type LogType int

const (
	INFO    LogType = iota // INFO log type, represented by 0.
	WARNING                // WARNING log type, represented by 1.
	ERROR                  // ERROR log type, represented by 2.
	DEBUG                  // DEBUG log type, represented by 3.
)

const (
	Reset   = "\033[0m"  // Resets the text color to default.
	Black   = "\033[30m" // Black color.
	Red     = "\033[31m" // Red color.
	Green   = "\033[32m" // Green color.
	Yellow  = "\033[33m" // Yellow color.
	Blue    = "\033[34m" // Blue color.
	Magenta = "\033[35m" // Magenta color.
	Cyan    = "\033[36m" // Cyan color.
	White   = "\033[37m" // White color.
)
