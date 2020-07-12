package youtube

import (
	"fmt"
	"os"
	"tinyfilter/dev/command/reload"
	"tinyfilter/dev/etc"
	"tinyfilter/dev/util"
)

// Reload TinyProxy configuration
func Exec(isOn bool) error {

	fmt.Println("Turning Youtube access", isOn)

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
