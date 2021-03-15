// Package constants contains some useful constants / helper functions
package constants

import (
	"log"
	"sync"
)

// RpiLibraryLogger is a logger for the rpi-radio-alarm-library
var (
	RpiLibraryLogger *log.Logger
	once             sync.Once
)

// GetLogger returns the Logger defined for the library
func GetLogger() *log.Logger {
	once.Do(func() {
		RpiLibraryLogger = log.New(log.Writer(), "[rpi-radio-alarm-library] \t", log.Ldate|log.Ltime|log.Lmsgprefix)
	})
	return RpiLibraryLogger
}
