package log

import (
	"time"
	"fmt"
)

// loggers holds the loggers to write to
var loggers = make([]Logger, 0)

// AddLogger adds the specified logger to be called when writing a log entry
func AddLogger(logger Logger) {
	loggers = append(loggers, logger)
}

// AddInfo adds a generic message of type info
func AddInfo(msg string, args ...interface{}) {
	Add(Info, msg, args...)
}

// Add a log message
func Add(logtype string, msg string, args ...interface{}) {
	newargs := make([]interface{}, 2)
	newargs[0] = time.Now().UTC()
	newargs[1] = logtype
	newargs = append(newargs, args...)
	
	msg = "%v - %v - " + msg
	for _,l := range loggers {
		l.Add(msg, newargs...)
	}
}