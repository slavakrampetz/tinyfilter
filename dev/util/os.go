package util

import (
	"runtime"
)

var (
	isWindows = false
	isFreeBsd = false
)

func init() {
	switch runtime.GOOS {
	case "freebsd" :
		isFreeBsd = true
		isWindows = false
		break
	case "windows":
		isWindows = true
		isFreeBsd = false
		break
	default:
		isWindows = true
		isFreeBsd = false
	}
}

func IsFreeBSD() bool {
	return isFreeBsd
}

func IsWindows() bool {
	return isWindows
}
