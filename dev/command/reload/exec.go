package reload

import (
	x "tinyfilter/dev/util/exec"
)

// Reload TinyProxy configuration
func Exec() error {

	err, _ := x.Exec(false, "service", "tinyproxy", "restart")
	if err != nil {
		return err
	}

	return nil
}
