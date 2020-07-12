package util

import (
	"runtime"
)

func IsFreeBSD() bool {
	return runtime.GOOS == "freebsd"
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}
