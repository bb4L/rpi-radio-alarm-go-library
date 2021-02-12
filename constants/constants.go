package constants

import (
	"log"
	"sync"
)

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
