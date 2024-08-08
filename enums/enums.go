package logify_enums

// LogType is an enumerated type for the different types of logs.
type LogType int

const (
	INFO    LogType = iota // LogType for common messages.
	WARNING                // LogType for warning messages.
	ERROR                  // LogType for error messages.
	DEBUG                  // LogType for debug messages.
)

// LogTypeParser converts LogType to a readable string.
func (env LogType) LogTypeParser() string {
	switch env {
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

// LoggerMode is an enumerated type for the different logging modes.
type LoggerMode int

const (
	LogModeDev     LoggerMode = iota // Developer mode, prints all logs but does not save them.
	LogModeRelease                   // Release mode, prints all logs except DEBUG.
	LogModeVerbose                   // Verbose mode, print all logs.
	LogModeHidden                    // Hidden mode, only prints INFO and ERROR logs.
)

// loggerModeParser converts LoggerMode to a readable string.
func (env LoggerMode) LoggerModeParser() string {
	switch env {
	case LogModeDev:
		return "DEVELOPER"
	case LogModeRelease:
		return "DEPLOYMENT"
	case LogModeVerbose:
		return "VERBOSE"
	case LogModeHidden:
		return "DISCRETE"
	default:
		return "UNKNOWN"
	}
}
