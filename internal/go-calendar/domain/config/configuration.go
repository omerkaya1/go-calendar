package config

import (
	"path"

	"github.com/omerkaya1/go-calendar/internal/go-calendar/domain/errors"
	"github.com/spf13/viper"
)

type (
	// Config is a struct that holds configuration parameters for the go-calendar app
	Config struct {
		Server     Server    `json:"server" yaml:"server" toml:"server"`
		DB         DB        `json:"db" yaml:"db" toml:"db"`
		Queue      QueueConf `json:"queue" yaml:"queue" toml:"queue"`
		Prometheus PromConf  `json:"prometheus" yaml:"prometheus" toml:"prometheus"`
	}
	// Server is a configuration struct for the server part of the app
	Server struct {
		Host  string `json:"host" yaml:"host" toml:"host"`
		Port  string `json:"port" yaml:"port" toml:"port"`
		Level int    `json:"level" yaml:"level" toml:"level"`
	}
	// DB is a configuration struct for the DB part of the app
	DB struct {
		Host     string `json:"host" yaml:"host" toml:"host"`
		Port     string `json:"port" yaml:"port" toml:"port"`
		Password string `json:"password" yaml:"password" toml:"password"`
		Name     string `json:"name" yaml:"name" toml:"name"`
		User     string `json:"user" yaml:"user" toml:"user"`
		SSL      string `json:"ssl" yaml:"ssl" toml:"ssl"`
	}
	// QueueConf is a configuration struct for the message queue part of the app
	QueueConf struct {
		Host     string `json:"host" yaml:"host" toml:"host"`
		Port     string `json:"port" yaml:"port" toml:"port"`
		User     string `json:"user" yaml:"user" toml:"user"`
		Password string `json:"password" yaml:"password" toml:"password"`
		Interval string `json:"interval" yaml:"interval" toml:"interval"`
		Name     string `json:"name" yaml:"name" toml:"name"`
	}
	// PromConf is a configuration struct for the monitoring part of the app
	PromConf struct {
		Host string `json:"host" yaml:"host" toml:"host"`
		Port string `json:"port" yaml:"port" toml:"port"`
	}
)

// Verify checks the configuration for the server
func (s Server) Verify() bool {
	return s.Host != "" || s.Port != ""
}

// Verify checks the configuration for the DB
func (qc DB) Verify() bool {
	return qc.Host != "" || qc.Port != "" || qc.User != "" || qc.Password != "" || qc.Name != "" || qc.SSL != ""
}

// Verify checks the configuration for the message queue service
func (qc QueueConf) Verify() bool {
	return qc.Host != "" || qc.Port != "" || qc.User != "" || qc.Password != "" || qc.Name != "" || qc.Interval != ""
}

// Verify checks the configuration for the monitoring service
func (s PromConf) Verify() bool {
	return s.Host != "" || s.Port != ""
}

// InitConf initialises configuration
func InitConfig(cfgPath string) (*Config, error) {
	viper.SetConfigFile(cfgPath)

	cfgFileExt := path.Ext(cfgPath)
	if cfgFileExt == "" {
		return nil, errors.ErrCorruptConfigFileExtension
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
