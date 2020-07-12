package reload

import (
	"testing"
	"tinyfilter/dev/util"
)

// Reload TinyProxy configuration
func TestExec(t *testing.T) {

	if !util.IsFreeBSD() {
		t.Skip("this test is for FreeBSD only")
	}

	err := Exec()
	if err != nil {
		t.Errorf("Failed: %s", err)
	}

}
