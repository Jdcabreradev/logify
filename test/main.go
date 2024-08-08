package main

import (
	"time"

	"github.com/Jdcabreradev/logify/v3"
	logify_enums "github.com/Jdcabreradev/logify/v3/enums"
)

func main() {
	// Create a logger that will save logs to a new file with a timestamp
	logger := logify.New("MyModule", "MyConsumer", "./logs", logify_enums.LogModeRelease)
	defer logger.Close() // Ensure the log file is closed when done

	logger.Log(logify_enums.INFO, "This is an informational message.")

	// Simulate application execution
	time.Sleep(2 * time.Second)

	//End of file
	logger.Log(logify_enums.ERROR, "This is an error message.")
}
