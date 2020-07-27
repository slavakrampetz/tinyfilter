package youtube

import (
	"tinyfilter/dev/command/reload"
	"tinyfilter/dev/etc"
	"tinyfilter/dev/log"
)

// Reload TinyProxy configuration
func Exec(isOn bool) error {

	log.Inf("Turning Youtube access", isOn)
	if err := etc.Config.Read(); err != nil {
		return err
	}

	if err := reLink(isOn); err != nil {
		return err
	}

	// Reload TinyProxy
	if err := reload.Exec(); err != nil {
		return err
	}

	return nil
}

// Reload TinyProxy configuration
func ExecRead() (Status, error) {

	log.Inf("Reading Youtube access")
	if err := etc.Config.Read(); err != nil {
		return StatusUnknown, err
	}

	return readLink()
}
