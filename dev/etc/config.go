package etc

import (
	"errors"
	tml "github.com/pelletier/go-toml"
	"os"
	"tinyfilter/dev/util"
)

type ConfigData struct {
	TinyProxy struct {
		Root   string
		Filter struct {
			Filename   string
			Default    string
			Restricted string
		}
	}
}

var (
	Config     ConfigData
	configPath string
)

func init() {
	Config = ConfigData{}
	configPath = GetPath()
}

func (c *ConfigData) Read() error {

	// No config
	if !util.IsFileReadable(configPath) {
		return errors.New("File " + configPath + " not readable")
	}

	doc, err := tml.LoadFile(configPath)
	if err != nil {
		return err
	}

	cfg := &ConfigData{}
	err = doc.Unmarshal(cfg)
	if err != nil {
		return err
	}

	err = cfg.Validate()
	if err != nil {
		return err
	}

	c.TinyProxy.Root = cfg.TinyProxy.Root
	c.TinyProxy.Filter.Filename = cfg.TinyProxy.Filter.Filename
	c.TinyProxy.Filter.Default = cfg.TinyProxy.Filter.Default
	c.TinyProxy.Filter.Restricted = cfg.TinyProxy.Filter.Restricted
	return nil
}

// Validate

func (c *ConfigData) Validate() error {

	if !util.IsDir(c.TinyProxy.Root) {
		return errors.New("Path " + c.TinyProxy.Root + " is not a directory or cannot be accessed")
	}

	fp := util.PathJoinSafe(c.TinyProxy.Root, c.TinyProxy.Filter.Filename)
	if !util.IsWindows() {
		if !util.IsLink(fp) {
			return errors.New("Path " + fp + " is not a readable symbolic link")
		}
	} else {
		if !util.IsFile(fp) {
			return errors.New("Path " + fp + " is not a readable file")
		}
	}

	fp = util.PathJoinSafe(c.TinyProxy.Root, c.TinyProxy.Filter.Default)
	if !util.IsFile(fp) {
		return errors.New("Default " + fp + " is not a readable file")
	}

	fp = util.PathJoinSafe(c.TinyProxy.Root, c.TinyProxy.Filter.Restricted)
	if !util.IsFile(fp) {
		return errors.New("Restricted " + fp + " is not a readable file")
	}

	return nil
}

func GetPath() string {
	if configPath == "" {
		configPath = util.PathReplaceExt(os.Args[0], ConfigExt)
	}
	return configPath
}
