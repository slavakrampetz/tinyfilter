package etc

import (
	"errors"
	"fmt"
	tml "github.com/pelletier/go-toml"
	"os"
	"tinyfilter/dev/util"
)

type ConfigData struct {

	// Config of TinyProxy ipc
	TinyProxy struct {
		Root   string
		Filter struct {
			Filename   string
			Default    string
			Restricted string
		}
	}

	// Web auth
	Auth struct {
		Type string
		Key  string
	}
}

var (
	Config       ConfigData
	configPath   string
	configIsRead bool
)

func init() {
	Config = ConfigData{}
	configPath = GetPath()
	configIsRead = false
}

func (c *ConfigData) Read() error {

	if configIsRead {
		return nil
	}

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
	configIsRead = true

	err = cfg.Validate()
	if err != nil {
		return err
	}

	c.TinyProxy.Root = cfg.TinyProxy.Root

	c.TinyProxy.Filter.Filename = cfg.TinyProxy.Filter.Filename
	c.TinyProxy.Filter.Default = cfg.TinyProxy.Filter.Default
	c.TinyProxy.Filter.Restricted = cfg.TinyProxy.Filter.Restricted

	c.Auth.Type = cfg.Auth.Type
	c.Auth.Key = cfg.Auth.Key
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

	// Auth key
	if c.Auth.Key == "" {
		return errors.New("Auth.Key not defined in config")
	}
	if c.Auth.Type == "" {
		c.Auth.Type = "query"
	} else {
		switch c.Auth.Type {
		case "query":
		case "header":
			break
		default:
			return fmt.Errorf("Auth.Type value is invalid: %s", c.Auth.Type)
		}
	}

	return nil
}

func GetPath() string {
	if configPath == "" {
		configPath = util.PathReplaceExt(os.Args[0], ConfigExt)
	}
	return configPath
}

func SetPath(p string) {
	configPath = p
}
