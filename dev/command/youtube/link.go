package youtube

import (
	"errors"
	"os"
	"tinyfilter/dev/etc"
	"tinyfilter/dev/util"
)

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
		source = etc.Config.TinyProxy.Filter.Default
	} else {
		source = etc.Config.TinyProxy.Filter.Restricted
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
