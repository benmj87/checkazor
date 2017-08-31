package main

import (
	"github.com/benmj87/checkazor/src/log"
)

// main holds the entry point to the application
func main() {
	log.AddLogger(log.NewConsoleLogger())
	log.AddInfo("Starting application %v at %v", "foo", "too")
}