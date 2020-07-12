package util

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

// IsDir check is path is exists and it's a directory
//noinspection GoUnusedExportedFunction
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
	return stat.Mode().IsRegular()
}

// IsLink check is path is exists and it's a symlink
func IsLink(path string) bool {
	stat, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}
	return stat.Mode()&os.ModeSymlink == os.ModeSymlink
}

//noinspection GoUnusedExportedFunction
func IsFileReadable(path string) bool {
	return IsFile(path) && isFileReadable(path)
}

//noinspection GoUnusedExportedFunction
func IsFileWriteable(path string) bool {
	return IsFile(path) && isFileWriteable(path)
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

func PathReplaceExt(path string, newExt string) string {
	var cleaned string
	ext := filepath.Ext(path)
	ext = strings.TrimSpace(ext)
	if len(ext) == 0 {
		cleaned = path
	} else {
		cleaned = strings.TrimSuffix(path, ext)
	}
	return cleaned + newExt
}

func pathSafePrefix(p string) string {
	if !IsWindows() {
		return ""
	}
	if strings.HasPrefix(p, "//") {
		return "/"
	}
	if strings.HasPrefix(p, "\\\\") {
		return "\\"
	}
	return ""
}

func PathDirSafe(p string) string {
	prefix := pathSafePrefix(p)
	return prefix + path.Dir(p)
}

func PathJoinSafe(parts ...string) string {
	prefix := ""
	if len(parts) > 0 {
		prefix = pathSafePrefix(parts[0])
	}
	res := path.Join(parts...)
	return prefix + res
}
