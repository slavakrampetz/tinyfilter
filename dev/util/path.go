package util

import (
	"os"
)

// IsDir check is path is exists and it's a directory
func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	if !stat.Mode().IsDir() {
		return false
	}
	return true
}

// IsFile check is path is exists and it's a directory
func IsFile(path string) bool {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	if !stat.Mode().IsRegular() {
		return false
	}

	return true
}

func IsFileReadable(path string) bool {
	return IsFile(path) && isFileReadable(path)
}

func IsFileWriteable(path string) bool {
	return IsFileReadable(path) && isFileWriteable(path)
}


func isFileReadable(path string) bool {
	f, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return false
	}
	_ = f.Close()
	return true
}

func isFileWriteable(path string) bool {
	f, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		return false
	}
	_ = f.Close()
	return true
}