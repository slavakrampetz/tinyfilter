package youtube

import (
	"tinyfilter/dev/command/reload"
	"tinyfilter/dev/etc"
	"tinyfilter/dev/log"
)

// Reload TinyProxy configuration
func Exec(isOn bool, minutes int) error {

	log.Inf("Turning Youtube access", isOn, "for", minutes, "min")
	if err := etc.Config.Read(); err != nil {
		return err
	}

	if err := linkRecreate(isOn); err != nil {
		return err
	}

	// Reload TinyProxy
	if err := reload.Exec(); err != nil {
		return err
	}

	// Update scheduler
	if isOn {
		autostopSet(minutes)
	} else {
		autostopCancel()
	}

	return nil
}

// Reload TinyProxy configuration
func ExecRead() (Status, error) {

	log.Inf("Reading Youtube access")
	if err := etc.Config.Read(); err != nil {
		return StatusUnknown, err
	}

	return linkRead()
}
