package log

import (
	"fmt"
)

func Out(args ...interface{}) {
	message := fmt.Sprintln(args...)
	printNonColoredMessage(Info, message)
}

func Eol() {
	printNonColoredMessage(Info, "")
}

// TODO:

// Ds prints a debug separator
// Only shown if DebugMode is set to true
func SEP(l Level, args ...interface{}) {
	if !isDebugMode() {
		return
	}
	message := formatMessage(args...)
	message = formatSeparator(message, "-", 80)
	printMessage(Debug, message)
}

