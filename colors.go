package logify

// LogColor holds the colors for each log type.
var LogColor = map[LogType]string{
	INFO:    Green,
	WARNING: Yellow,
	ERROR:   Red,
	DEBUG:   Blue,
}

// SetColor allows changing the color for a specific log type.
func SetColor(logType LogType, color string) {
	if _, valid := map[string]bool{
		Black: true, Red: true, Green: true, Yellow: true,
		Blue: true, Magenta: true, Cyan: true, White: true,
	}[color]; valid {
		LogColor[logType] = color
	}
}
