// Package logging contains some logging utils
package logging

import (
	"io"
	"log"
)

const commonPrefix = "[rpi-radio-alarm-library] "

var prefixes = log.Ldate | log.Ltime | log.Lmsgprefix

// GetLogger returns a logger with given addional prefix
func GetLogger(place string, output io.Writer) *log.Logger {
	return log.New(output, commonPrefix+"["+place+"] ", prefixes)
}
