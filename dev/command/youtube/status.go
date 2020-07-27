package youtube

import (
	"time"
	"tinyfilter/dev/command/reload"
	"tinyfilter/dev/etc"
	"tinyfilter/dev/log"
)

type Status uint8

const (
	StatusUnknown = Status(0)
	StatusOn      = Status(1)
	StatusOff     = Status(2)
)

func (s Status) String() string {
	switch s {
	case StatusOn:
		return "On"
	case StatusOff:
		return "Off"
	}
	return "Unknown"
}
