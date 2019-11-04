package config

import (
	"path"

	"github.com/omerkaya1/go-calendar/internal/watcher/errors"
	"github.com/spf13/viper"
)

// Config .
type Config struct {
	Host  string    `json:"host" yaml:"host" toml:"host"`
	Port  string    `json:"port" yaml:"port" toml:"port"`
	Level int       `json:"level" yaml:"level" toml:"level"`
	DB    DBConf    `json:"db" yaml:"db" toml:"db"`
	Queue QueueConf `json:"queue" yaml:"queue" toml:"queue"`
}

// DBConf .
type DBConf struct {
	Host     string `json:"host" yaml:"host" toml:"host"`
	Port     string `json:"port" yaml:"port" toml:"port"`
	Password string `json:"password" yaml:"password" toml:"password"`
	Name     string `json:"name" yaml:"name" toml:"name"`
	User     string `json:"user" yaml:"user" toml:"user"`
	SSL      string `json:"ssl" yaml:"ssl" toml:"ssl"`
}

// QueueConf .
type QueueConf struct {
	Host     string `json:"host" yaml:"host" toml:"host"`
	Port     string `json:"port" yaml:"port" toml:"port"`
	User     string `json:"user" yaml:"user" toml:"user"`
	Password string `json:"password" yaml:"password" toml:"password"`
	Interval string `json:"interval" yaml:"interval" toml:"interval"`
	Name     string `json:"name" yaml:"name" toml:"name"`
}

// InitConf .
func InitConfig(cfgPath string) (*Config, error) {
	viper.SetConfigFile(cfgPath)

	cfgFileExt := path.Ext(cfgPath)
	if cfgFileExt == "" {
		return nil, errors.ErrBadConfigFile
	}
	viper.SetConfigType(cfgFileExt[1:])

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
