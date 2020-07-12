package util

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestPathReplaceExt(t *testing.T) {

	var tests = []struct {
		s string
		r string
	}{
		{"filename.ext", "filename.config"},
		{"filename", "filename.config"},
		{"t:\\path\\to\\file\\filename.ext", "t:\\path\\to\\file\\filename.config"},
		{"t:/path/to/file/filename.ext", "t:/path/to/file/filename.config"},
		{"/mnt/filename.pid", "/mnt/filename.config"},
	}

	// 1. Test getting extension
	exe := os.Args[0]
	base := filepath.Base(exe)
	ext := filepath.Ext(base)
	if IsWindows() {
		if ext != ".exe" {
			t.Errorf("got %s want .exe", ext)
		}
	} else {
		if ext != "" {
			t.Errorf("got %s want <empty>", ext)
		}
	}

	for idx, tt := range tests {
		test := fmt.Sprintf("#%d", idx)
		t.Run(test, func(t *testing.T) {
			res := PathReplaceExt(tt.s, ".config")
			if res != tt.r {
				t.Errorf("got '%s', want '%s'", res, tt.r)
			}
		})
	}

}
