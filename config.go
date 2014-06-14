package main

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os/user"
	"path"
)

// File that user information such as client id and
// api keys will be written to.
var (
	configFile string
	config     Config
)

func init() {
	var u *user.User
	if u, _ = user.Current(); u == nil {
		panic("user.Current is nil")
	}
	configFile = path.Join(u.HomeDir, ".fish")
}

// Config is the Configuration settings.
type Config struct {
	Conf     conf
	Defaults defaults
}

// Parsed configuration settings.
type conf struct {
	ClientID string `toml:"client_id"`
	APIKey   string `toml:"api_key"`
}

type defaults struct {
	Size           string `toml:"size"`
	Region         string `toml:"region"`
	Image          int    `toml:"image"`
	PrivateNetwork bool   `toml:"private_network"`
	Backups        bool   `toml:"backups_enabled"`
}

// Loads the configuration from the filepath. If there's any
// issue an error will be returned.
func loadConfig() error {
	// The file exists, let's parse it
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	_, err = toml.Decode(string(bytes), &config)
	if err != nil {
		return err
	}
	return nil
}
