package log

import (
	"fmt"
	"github.com/sanity-io/litter"
)

// Dbg prints a debug message if DebugMode is set to true
func Dbg(args ...interface{}) {
	if !isDebugMode() {
		return
	}
	message := formatMessage(args...)
	printMessage(Debug, message)
}

// DbgIf prints a debug message if DebugMode is set to true and condition met
func DbgIf(condition bool, args ...interface{}) {
	if !condition || !isDebugMode() {
		return
	}
	Dbg(args...)
}

// DbgErr prints a debug message if DebugMode is set to true and error is not nil
func DbgErr(err error, args ...interface{}) {
	if err == nil || !isDebugMode() {
		return
	}
	Err(err.Error())
	Dbg(args...)
}

// DbgF prints a debug message with a format and arguments if DebugMode is set to true
func DbgF(format string, args ...interface{}) {
	if !isDebugMode() {
		return
	}
	message := fmt.Sprintf(format, args...)
	Dbg(message)
}

// DbgFIf prints a debug message with a format and arguments if DebugMode is set to true and condition met
func DbgFIf(condition bool, format string, args ...interface{}) {
	if !condition || !isDebugMode() {
		return
	}
	DbgF(format, args...)
}

// DbgFErr prints a debug message with a format and arguments if DebugMode is set to true and error happens
func DbgFErr(err error, format string, args ...interface{}) {
	if err == nil || !isDebugMode() {
		return
	}
	Err(err.Error())
	DbgF(format, args...)
}

// DbgDump dumps the argument as a debug message with an optional prefix
func DbgDump(prefix string, arg interface{}) {
	if !isDebugMode() {
		return
	}
	message := litter.Sdump(arg)
	if prefix != "" {
		message = formatMessage(prefix, arg)
	}
	printMessage(Debug, message)
}

// DbgDumpIf dumps the argument as a debug message with an optional prefix if condition met
func DbgDumpIf(condition bool, prefix string, arg interface{}) {
	if !condition || !isDebugMode() {
		return
	}
	DbgDump(prefix, arg)
}

// DbgDumpErr dumps the argument as a debug message with an optional prefix if error happens
func DbgDumpErr(err error, prefix string, arg interface{}) {
	if err == nil || !isDebugMode() {
		return
	}
	Err(err.Error())
	DbgDump(prefix, arg)
}


// DbgTrace prints an error message with the stacktrace of err to stderr
func DbgTrace(err error) {
	if err == nil || !isDebugMode() {
		return
	}
	message := formatMessage(GetTrace(err))
	printMessage(Debug, message)
}

