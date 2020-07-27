package youtube

import (
	"time"
	"tinyfilter/dev/log"
)

var (
	// Is Youtube allowed for now
	until *time.Time
)

func init() {
	until = nil
}

// TODO: See timers at
// https://gobyexample.ru/timeouts.html

func autostopCancel() {
	log.Inf("autostopCancel: not implemented yet. Current:", until)
}

func autostopSet(minutes int) {
	log.Inf("autostopSet: not implemented yet, minutes", minutes)
}
