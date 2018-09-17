package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Mode :
var Mode = "default"

// Debug is false by default
var Debug = false

// Version : default version
var Version = "v0.0.1"

// Configurations : manage all configurations
type Configurations struct {
	Version                   string   `yaml:"version"`
	Location                  string   `yaml:"location"`
	Debug                     bool     `yaml:"debug"`
	AvailableEnvironmentFlags []string `yaml:"available_environment_flags"`
	// server conf
	Server *Server `yaml:"server"`
	// redis conf
	RedisConf *RedisConf `yaml:"redisconf"`
}

// Init : init all configurations
func (c *Configurations) Init() error {
	c.Version = Version
	c.Location = "Local"
	c.Debug = Debug

	// server
	c.Server = &Server{}
	c.Server.Init()

	// redis init
	c.RedisConf = &RedisConf{}
	c.RedisConf.Init()

	return nil
}

// GlobalConfigurations :
var GlobalConfigurations = &Configurations{}

// ReadFrom : read from file
func (c *Configurations) ReadFrom(path string) error {
	c.Init()

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &GlobalConfigurations)
	if err != nil {
		return err
	}
	if GlobalConfigurations.Version != Version {
		return fmt.Errorf("The version of base file is different with the config file")
	}
	return nil
}
