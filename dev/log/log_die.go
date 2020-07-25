package log

import (
	"fmt"
)

// Die logs a fatal error message to stdout and exits the program with exit code 1
func Die(args ...interface{}) {
	message := formatMessage(args...)
	printMessage(FatalL, message)
	OsExit(1)
}

// DieIf logs a fatal error message to stdout and exits the program with exit code 1
func DieIf(condition bool, args ...interface{}) {
	if !condition {
		return
	}
	Die(args...)
}

// DieErr logs a fatal error message to stdout and exits the program with exit code 1
func DieErr(err error, args ...interface{}) {
	if err == nil {
		return
	}
	printMessage(FatalL, err.Error())
	Die(args...)
}

// DieF prints a fatal message with a format and arguments
func DieF(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	Die(message)
}

// DieFIf prints a fatal message with a format and arguments
func DieFIf(condition bool, format string, args ...interface{}) {
	if !condition {
		return
	}
	DieF(format, args...)
}

// DieExit checks if the error is not nil and if that's the case,
// it will print a fatal message and exits the program with exit code 1.
//
// If DebugMode is enabled a stack trace will also be printed to stderr
func DieExit(err error, code int) {
	if err == nil {
		return
	}
	if code == 0 {
		code = 1
	}
	var message string
	if isDebugMode() {
		message = formatMessage(GetTrace(err))
	} else {
		message = err.Error()
	}
	printMessage(FatalL, message)
	OsExit(code)
}
