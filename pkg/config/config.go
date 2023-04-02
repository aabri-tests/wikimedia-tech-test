package config

import (
	"github.com/jinzhu/configor"
	"github.com/pkg/errors"
)

const (
	defaultMongoDBTimeout = 5
)

type Config struct {
	Name   string `yaml:"name"`
	Logger Logger `yaml:"logger"`
	Server Server `yaml:"server"`
	Cache  Cache  `yaml:"cache"`
}

// New creates a new instance of the Config structure and loads configuration from a config.yml file using the configor package
func New() (*Config, error) {
	config := &Config{}
	errLoad := configor.Load(config, "config.yml")
	if errLoad != nil {
		return nil, errors.Wrap(errLoad, "failed to load configuration file")
	}
	return config, nil
}
