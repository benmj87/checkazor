package log

import (
	"fmt"
)

// ConsoleLogger writes logs out to the console
type ConsoleLogger struct {
}

// NewConsoleLogger returns an instance of the console logger
func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

// Add a log message
func (c *ConsoleLogger) Add(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}