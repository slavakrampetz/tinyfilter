package youtube

import (
	"io/ioutil"
	"path/filepath"
	"testing"
	"tinyfilter/dev/etc"
	"tinyfilter/dev/util"
)

// Reload TinyProxy configuration
func TestExec(t *testing.T) {

	// Prepare
	if !util.IsFreeBSD() {
		t.Skip("this test is for FreeBSD only")
	}

	from, line := util.DebugFrom() // //data/golang/src/tinyfilter/dev/etc/config_test.go
	if line < 0 {
		panic("cannot get path to current source file")
	}
	dir := util.PathDirSafe(from)

	configPath := util.PathJoinSafe(dir, "../../../test/command.youtube/test.config")
	etc.SetPath(configPath)


	// test readlink
	save := etc.Config.TinyProxy.Filter.Filename
	etc.Config.TinyProxy.Filter.Filename = "non-existing-file"
	status, err := readLink()
	if err == nil {
		t.Errorf("Expect error on read non-existing file, got mil")
	}
	if status != StatusUnknown {
		t.Errorf("Expect %v on read non-existing file, got %v", StatusUnknown, status)
		return
	}
	etc.Config.TinyProxy.Filter.Filename = save

	// Exec
	var text string

	err = Exec(true)
	if err != nil {
		t.Errorf("Failed: %s", err)
		return
	}
	err, text = readTestDataByLink()
	if err != nil {
		t.Errorf("Cannot read link: %s", err)
	}
	if text != "restricted" {
		t.Errorf("Expect 'restricted', got '%s'", text)
		return
	}

	status, err = ExecRead()
	if err != nil {
		t.Errorf("Failed: %s", err)
		return
	}
	if status != StatusOff {
		t.Errorf("Expect %v, got %v", StatusOff, status)
		return
	}

	err = Exec(false)
	if err != nil {
		t.Errorf("Failed: %s", err)
		return
	}
	err, text = readTestDataByLink()
	if err != nil {
		t.Errorf("Cannot read link: %s", err)
	}
	if text != "default" {
		t.Errorf("Expect 'default', got '%s'", text)
		return
	}

	status, err = ExecRead()
	if err != nil {
		t.Errorf("Failed: %s", err)
		return
	}
	if status != StatusOn {
		t.Errorf("Expect %v, got %v", StatusOn, status)
		return
	}
}

func readTestDataByLink() (error, string) {
	name := etc.Config.TinyProxy.Filter.Filename
	p := util.PathJoinSafe(etc.Config.TinyProxy.Root, name)

	target, err := filepath.EvalSymlinks(p)
	if err != nil {
		return err, ""
	}
	data, err := ioutil.ReadFile(target)
	if err != nil {
		return err, ""
	}
	return nil, string(data)
}
