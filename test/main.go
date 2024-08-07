package main

import (
	"time"

	"github.com/Jdcabreradev/logify"
)

func main() {
	// Create a logger that will save logs to a new file with a timestamp
	logger := logify.NewLogger("MyModule", "MyConsumer", "development", "./logs/")
	defer logger.Close() // Ensure the log file is closed when done

	logger.Log(logify.INFO, "This is an informational message.")

	// Simulate application execution
	time.Sleep(2 * time.Second)

	//End of file
	logger.Log(logify.ERROR, "This is an error message.")
}
