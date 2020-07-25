package log

import (
	"fmt"
	"github.com/sanity-io/litter"
)

// Err prints an error message to stderr
func Err(args ...interface{}) {
	message := formatMessage(args...)
	printMessage(Error, message)
}

// ErrIf prints an error message to stderr if condition met
func ErrIf(condition bool, args ...interface{}) {
	if !condition {
		return
	}
	Err(args...)
}

// ErrIfErr prints an error message to stderr if error detected
func ErrIfErr(err error, args ...interface{}) {
	if err == nil {
		return
	}
	Err(err.Error())
	Err(args...)
}

// ErrF prints an error message with a format and arguments
func ErrF(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	Err(message)
}

// ErrFIf prints an error message with a format and arguments if condition met
func ErrFIf(condition bool, format string, args ...interface{}) {
	if !condition {
		return
	}
	ErrF(format, args...)
}

// ErrFIfErr prints an error message with a format and arguments if error exist
func ErrFIfErr(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}
	Err(err.Error())
	ErrF(format, args...)
}

// ErrDump dumps the argument as an err message with an optional prefix to stderr
func ErrDump(prefix string, arg interface{}) {
	message := litter.Sdump(arg)
	if prefix != "" {
		Err(prefix, message)
	} else {
		Err(message)
	}
}

// ErrDumpIf dumps the argument as an err message with an optional prefix to stderr
func ErrDumpIf(condition bool, prefix string, arg interface{}) {
	if !condition {
		return
	}
	ErrDump(prefix, arg)
}

// ErrDumpErr dumps the argument as an err message with an optional prefix to stderr
func ErrDumpErr(err error, prefix string, arg interface{}) {
	if err == nil {
		return
	}
	Err(err.Error())
	ErrDump(prefix, arg)
}

// ErrTrace prints an error message with the stacktrace of err to stderr
func ErrTrace(err error) {
	if err == nil {
		return
	}
	message := formatMessage(GetTrace(err))
	printMessage(Error, message)
}
