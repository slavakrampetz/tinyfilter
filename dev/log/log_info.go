package log

import (
	"fmt"
	"github.com/sanity-io/litter"
)

// Info prints an info message
func Inf(args ...interface{}) {
	message := formatMessage(args...)
	printMessage(Info, message)
}

// Info prints an info message if condition met
func InfIf(condition bool, args ...interface{}) {
	if !condition {
		return
	}
	Inf(args...)
}

// Info prints an info message if error not empty
func InfErr(err error, args ...interface{}) {
	if err == nil {
		return
	}
	Err(err)
	Inf(args...)
}

// InfF prints an info message with a format and arguments
func InfF(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	Inf(message)
}

// InfFIf prints an info message with a format and arguments if condition met
func InfFIf(condition bool, format string, args ...interface{}) {
	if !condition {
		return
	}
	InfF(format, args...)
}

// InfFErr prints an info message with a format and arguments if error not empty
func InfFErr(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}
	Err(err.Error())
	InfF(format, args...)
}

// InfDump dumps the argument as an info message with an optional prefix
func InfDump(prefix string, arg interface{}) {
	message := litter.Sdump(arg)
	if prefix != "" {
		Inf(prefix, message)
	} else {
		Inf(message)
	}
}

// InfDumpIf dumps the argument as an info message with an optional prefix if condition met
func InfDumpIf(condition bool, prefix string, arg interface{}) {
	if !condition {
		return
	}
	InfDump(prefix, arg)
}

// InfDumpErr dumps the argument as an info message with an optional prefix if error not empty
func InfDumpErr(err error, prefix string, arg interface{}) {
	if err == nil {
		return
	}
	Err(err)
	InfDump(prefix, arg)
}
