package log

import (
)

// Logger interface to add log messages
type Logger interface {
	// Add a message
	Add(msg string, args ...interface{})
}