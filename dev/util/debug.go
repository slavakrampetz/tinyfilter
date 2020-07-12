package util

import (
	"runtime"
)

func DebugFrom() (string, int) {
	_, fileName, fileLine, ok := runtime.Caller(1)
	if !ok {
		return "UNDEFINED", -1
	}
	return fileName, fileLine
}
