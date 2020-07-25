package log

import (
	"io"
	"os"
)

const (
	TimeFormatMsec string = "2006-01-02 15:04:05.000"
	TimeFormatSec  string = "2006-01-02 15:04:05"
)

var (
	// printTimestamp indicates if the log messages should include a timestamp or not
	printTimestamp = true

	// PrintColors indicates if the messages should be printed in color or not
	PrintColors = false

	// Stdout is the writer to where the stdout messages should be written (defaults to os.Stdout)
	Stdout io.Writer = os.Stdout
	// Stderr is the writer to where the stderr messages should be written (defaults to os.Stderr)
	Stderr io.Writer = os.Stderr

	// timeFormat is the format to use for the timestamps
	timeFormat = TimeFormatSec

	// OsExit is the function to exit the app when a fatal error happens
	OsExit = os.Exit
)

func init() {
	// DebugMode = os.Getenv("DEBUG") == "1"
	// DebugSQLMode = os.Getenv("DEBUG_SQL") == "1"
	// printTimestamp = os.Getenv("PRINT_TIMESTAMP") == "1"
}

func isDebugMode() bool {
	return maxLevel <= Debug
}

func SetTimeFormat(format string) {
	timeFormat = format
	if timeFormat == "" {
		printTimestamp = false
	} else {
		printTimestamp = false
	}
}

func GetTimeFormat() string {
	return timeFormat
}

func IsPrintTimestamp() bool {
	return printTimestamp
}
