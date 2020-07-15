package reload

import (
	"github.com/labstack/gommon/log"
	"tinyfilter/dev/util"
	x "tinyfilter/dev/util/exec"
)

// Reload TinyProxy configuration
func Exec() error {

	if util.IsWindows() {
		log.Warn("Cannot reload config under Windows. Skipping.")
		return nil
	}

	err, _ := x.Exec(false, "service", "tinyproxy", "restart")
	if err != nil {
		return err
	}

	return nil
}
