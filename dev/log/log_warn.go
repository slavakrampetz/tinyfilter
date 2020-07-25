package log

import (
	"fmt"
	"github.com/sanity-io/litter"
)

// Warn prints an warning message
func Warn(args ...interface{}) {
	message := formatMessage(args...)
	printMessage(Warning, message)
}

// WarnIf prints an warning message if condition met
func WarnIf(condition bool, args ...interface{}) {
	if !condition {
		return
	}
	Warn(args...)
}

// WarnIfErr prints an warning message if error exist
func WarnIfErr(err error, args ...interface{}) {
	if err == nil {
		return
	}
	Err(err.Error())
	Warn(args...)
}

// WarnF prints a warning message with a format and arguments
func WarnF(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	Warn(message)
}

// WarnFIf prints a warning message with a format and arguments if condition met
func WarnFIf(condition bool, format string, args ...interface{}) {
	if !condition {
		return
	}
	WarnF(format, args...)
}

// WarnFErr prints a warning message with a format and arguments if error detected
func WarnFErr(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}
	Err(err.Error())
	WarnF(format, args...)
}

// WarnDump dumps the argument as a warning message with an optional prefix
func WarnDump(prefix string, arg interface{}) {
	message := litter.Sdump(arg)
	if prefix != "" {
		Warn(prefix, message)
	} else {
		Warn(message)
	}
}

// WarnDumpIf dumps the argument as a warning message with an optional prefix if condition met
func WarnDumpIf(condition bool, prefix string, arg interface{}) {
	if !condition {
		return
	}
	WarnDump(prefix, arg)
}

// WarnDumpErr dumps the argument as a warning message with an optional prefix if error present
func WarnDumpErr(err error, prefix string, arg interface{}) {
	if err == nil {
		return
	}
	WarnDump(prefix, arg)
}
