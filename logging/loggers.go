// Package logging contains utils to get a logger with e standardized formatting
package logging

import (
	"io"
	"log"
	"strings"
)

const prefixes = log.Ldate | log.Ltime | log.Lmsgprefix

// GetLogger returns a logger with given additional prefix
func GetLogger(output io.Writer, places ...string) *log.Logger {
	return log.New(output, "["+strings.Join(places, ".")+"] ", prefixes)
}
