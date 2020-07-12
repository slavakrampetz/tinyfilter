package etc

import (
	"fmt"
	"os"
	"runtime"
	"testing"
	"tinyfilter/dev/util"
)

func init() {
	from, line := util.DebugFrom() // //data/golang/src/tinyfilter/dev/etc/config_test.go
	if line < 0 {
		panic("cannot get path to current source file")
	}
	dir := util.PathDirSafe(from)

	basename := runtime.GOOS
	configPath = util.PathJoinSafe(dir, "/../../test/configtest/"+basename+".config")
}

func Test1(t *testing.T) {

	exe := os.Args[0]
	fmt.Println("Exe:", exe)

	config := GetPath()
	fmt.Println("Config:", config)
}

func TestReadConfig(t *testing.T) {

	err := Config.Read()
	if err != nil {
		t.Error("Error reading config", GetPath(), err)
	}
}
