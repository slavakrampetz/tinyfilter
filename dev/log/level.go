package log

import (
	"strconv"
)

type Level uint8

const (
	Trace Level = 0
	Debug Level = 30
	Info  Level = 100

	// Problem notification levels
	Warning Level = 160
	Error   Level = 180
	FatalL  Level = 200
)

const (
	labelFatal = "fatal"
	labelErr = "error"
	labelDebug = "debug"
	labelWarn = "warning"
	labelInfo = "info"
	labelTrace = "trace"
)

var (
	allLevels = []Level{FatalL, Error, Warning, Info, Debug, Trace}
	maxLevel  = Warning

	prefixes map[uint8]string
)

func init() {
	prefixes = make(map[uint8]string)
}

func (l Level) String() string {
	switch l {
	case Trace:
		return labelTrace
	case Debug:
		return labelDebug
	case Info:
		return labelInfo
	case Warning:
		return labelWarn
	case Error:
		return labelErr
	case FatalL:
		return labelFatal
	}
	return "lv-" + strconv.FormatUint(uint64(l), 10)
}

func GetLevel() Level {
	return maxLevel
}

// Add custom level
func AddCustomLevel(level uint8, label string) {
	prefixes[level] = label
	postAddLevels()
}

// Bulk add custom levels
func AddCustomLevels(levels map[uint8]string) {
	for level, label := range levels {
		prefixes[level] = label
	}
	postAddLevels()
}

func postAddLevels() {
	var max uint8
	for level := range prefixes {
		if level < max {
			max = level
		}
	}
	var knownMax Level
	for _, known := range allLevels {
		if max < uint8(known) {
			knownMax = known
			break
		}
	}
	maxLevel = knownMax
}
