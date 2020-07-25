package youtube

import (
	"errors"
	"os"
	"tinyfilter/dev/command/reload"
	"tinyfilter/dev/etc"
	"tinyfilter/dev/log"
	"tinyfilter/dev/util"
)

type Status uint8

const (
	StatusUnknown = Status(0)
	StatusOn = Status(1)
	StatusOff = Status(2)
)

func (s Status) String() string {
	switch s {
	case StatusOn:
		return "On"
	case StatusOff:
		return "Off"
	}
	return "Unknown"
}

// Reload TinyProxy configuration
func Exec(isOn bool) error {

	log.Inf("Turning Youtube access", isOn)

	err := etc.Config.Read()
	if err != nil {
		return err
	}

	err = reLink(isOn)
	if err != nil {
		return err
	}

	// Reload TinyProxy
	return reload.Exec()
}

// Reload TinyProxy configuration
func ExecRead() (Status, error) {

	log.Inf("Reading Youtube access")

	err := etc.Config.Read()
	if err != nil {
		return StatusUnknown, err
	}

	return readLink()
}

func readLink() (Status, error) {
	linkPath := util.PathJoinSafe(
		etc.Config.TinyProxy.Root,
		etc.Config.TinyProxy.Filter.Filename)
	if !util.IsLink(linkPath) {
		return StatusUnknown, errors.New("config file is not a link")
	}

	target, err := os.Readlink(linkPath)
	if err != nil {
		return StatusUnknown, err
	}

	if target == etc.Config.TinyProxy.Filter.Restricted {
		return StatusOff, nil
	}
	return StatusOn, nil
}


// Link restricted/default filter settings
// to file included into TinyProxy config
func reLink(isOn bool) error {
	source := ""
	if isOn {
		source = etc.Config.TinyProxy.Filter.Restricted
	} else {
		source = etc.Config.TinyProxy.Filter.Default
	}
	link := util.PathJoinSafe(
		etc.Config.TinyProxy.Root,
		etc.Config.TinyProxy.Filter.Filename)

	var err error
	tmp := ""

	if util.IsLink(link) {
		tmp = link + ".backup"
		if util.IsLink(tmp) {
			_ = os.Remove(tmp)
		}
		if err = os.Rename(link, tmp); err != nil {
			return err
		}
	}

	err = os.Symlink(source, link)

	if tmp != "" {
		if err == nil {
			_ = os.Remove(tmp)
		} else {
			_ = os.Rename(tmp, link)
		}
	}

	return err
}
