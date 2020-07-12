package youtube

import (
	"testing"
)

// Reload TinyProxy configuration
func TestExec(t *testing.T) {

	// if !util.IsFreeBSD() {
	// 	t.Skip("this test is for FreeBSD only")
	// }

	var err error

	err = Exec(true)
	if err != nil {
		t.Errorf("Failed: %s", err)
	}

	err = Exec(false)
	if err != nil {
		t.Errorf("Failed: %s", err)
	}

}
