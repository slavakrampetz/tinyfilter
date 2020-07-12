package exec

import (
	"regexp"
	"strings"
	"testing"
	"tinyfilter/dev/util"
)

const (
	tinyproxyBin           = "tinyproxy"
	tinyproxyStatusPattern = `^` + tinyproxyBin + ` is running as pid \d+\.$`
	serviceStatusCommand   = "status"
)

func TestExecEcho1(t *testing.T) {
	err, out := Exec(true, "echo", "1")

	if err != nil {
		t.Errorf("exec 'echo 1' failed: %s", err)
		return
	}

	if strings.Trim(out, "\r\n") != "1" {
		t.Errorf("exec 'echo 1' failed: expect '1', got %s", out)
	}
}

func TestExecGo(t *testing.T) {
	err, _ := Exec(false, "go", "version")
	if err != nil {
		t.Errorf("exec 'go version' failed: %s", err)
		return
	}
}

func TestExecServiceFreeBsd(t *testing.T) {

	if !util.IsFreeBSD() {
		return
	}

	err, out := Exec(false, "service", tinyproxyBin, serviceStatusCommand)
	if err != nil {
		t.Errorf("exec 'service status' failed: %s", err)
		return
	}

	// tinyproxy is running as pid 38227.
	out = strings.Trim(out, "\r\n")

	if matched, err := regexp.MatchString(tinyproxyStatusPattern, out); !matched || err != nil {
		t.Errorf("exec 'service status' failed: %s", err)
	}
}
