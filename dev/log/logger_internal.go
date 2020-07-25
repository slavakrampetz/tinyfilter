package log

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
	"tinyfilter/dev/util"

	"github.com/fatih/color"
)

var (
	colors = map[Level]*color.Color{
		Debug: color.New(color.FgHiBlack),
		Info: color.New(color.FgHiGreen),
		Warning: color.New(color.FgHiYellow),
		Error: color.New(color.FgHiRed),
		FatalL: color.New(color.FgHiRed),
	}

	mutex sync.Mutex
)


func init() {
	color.NoColor = false
}

func formatMessage(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	msg = strings.TrimRight(msg, " \n\r")
	return msg
}

func formatSeparator(message string, separator string, length int) string {
	if message == "" {
		return strings.Repeat(separator, length)
	}
	prefix := strings.Repeat(separator, 4)
	suffixLength := length - len(message) - len(prefix) - 4
	suffix := ""
	if suffixLength > 0 {
		suffix = strings.Repeat(separator, suffixLength)
	}
	return prefix + "[ " + message + " ]" + suffix
}

func printMessage(level Level, message string) {

	label := level.String()

	if IsPrintTimestamp() {
		message = addTimestampToMessage(label, message)
	}

	if PrintColors && !util.IsWindows() {
		printColoredMessage(level, message)
	} else {
		printNonColoredMessage(level, message)
	}

}

func printNonColoredMessage(level Level, message string) {
	w := writerForLevel(level)
	_, _ = w.Write([]byte(message + "\n"))
}

func printColoredMessage(level Level, message string) {
	w := writerForLevel(level)
	if c, ok := colors[level]; ok {
		mutex.Lock()
		defer mutex.Unlock()
		c.EnableColor()
		_, _ = c.Fprint(w, message)
		_, _ = w.Write([]byte("\n"))
	} else {
		_, _ = w.Write([]byte(message + "\n"))
	}
}

func addTimestampToMessage(label string, message string) string {
	formattedTime := time.Now().Format(timeFormat)
	message = formattedTime + " " + label + "> " + message
	return message
}

func writerForLevel(level Level) io.Writer {
	if level < Error {
		return Stdout
	}
	return Stderr
}
