package logify_colors

import logify_enums "github.com/Jdcabreradev/logify/v3/enums"

const (
	Reset   = "\033[0m"  // Resets the text color to default.
	Black   = "\033[30m" // Black color.
	Red     = "\033[31m" // Red color.
	Green   = "\033[32m" // Green color.
	Yellow  = "\033[33m" // Yellow color.
	Blue    = "\033[34m" // Blue color.
	Magenta = "\033[35m" // Magenta.
	Cyan    = "\033[36m" // Cyan color.
	White   = "\033[37m" // White color.
)

// LogColor is a map that associates each log type with a specific color for console output.
var LogColor = map[logify_enums.LogType]string{
	logify_enums.INFO:    Green,  // The color green is used for INFO log types.
	logify_enums.WARNING: Yellow, // The color yellow is used for WARNING log types.
	logify_enums.ERROR:   Red,    // The color red is used for ERROR log types.
	logify_enums.DEBUG:   Blue,   // The color blue is used for DEBUG log types.
}

// SetColor enables changing the color associated with a specific log type.
func SetColor(logType logify_enums.LogType, color string) {
	validColors := []string{Black, Red, Green, Yellow, Blue, Magenta, Cyan, White}
	for _, c := range validColors {
		if c == color {
			LogColor[logType] = color
			return
		}
	}
}
