package logify

// LogColor holds the colors for each log type.
// LogColor is a map that associates each log type with a specific color for console output.
var LogColor = map[LogType]string{
	INFO:    Green,  // The color green is used for INFO log types.
	WARNING: Yellow, // The color yellow is used for WARNING log types.
	ERROR:   Red,    // The color red is used for ERROR log types.
	DEBUG:   Blue,   // The color blue is used for DEBUG log types.
}

// SetColor allows changing the color for a specific log type.
// SetColor enables changing the color associated with a specific log type.
func SetColor(logType LogType, color string) {
	// Map of valid colors. The color must be in this map to be considered valid.
	if _, valid := map[string]bool{
		Black:   true, // Valid black color.
		Red:     true, // Valid red color.
		Green:   true, // Valid green color.
		Yellow:  true, // Valid yellow color.
		Blue:    true, // Valid blue color.
		Magenta: true, // Valid magenta color.
		Cyan:    true, // Valid cyan color.
		White:   true, // Valid white color.
	}[color]; valid {
		// If the color is valid, update the color associated with the log type in LogColor.
		LogColor[logType] = color
	}
}
