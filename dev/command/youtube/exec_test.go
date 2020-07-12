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

	// Exec
	var err error
	var text string

	err = Exec(true)
	if err != nil {
		t.Errorf("Failed: %s", err)
		return
	}
	err, text = readLink()
	if err != nil {
		t.Errorf("Cannot read link: %s", err)
	}
	if text != "restricted" {
		t.Errorf("Expect 'restricted', got '%s'", text)
		return
	}

	err = Exec(false)
	if err != nil {
		t.Errorf("Failed: %s", err)
		return
	}
	err, text = readLink()
	if err != nil {
		t.Errorf("Cannot read link: %s", err)
	}
	if text != "default" {
		t.Errorf("Expect 'restricted', got '%s'", text)
		return
	}

}

func readLink() (error, string) {
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
