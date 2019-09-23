package conf

import (
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/spf13/viper"
	"path"
)

type Config struct {
	Host     string `json:"host" yaml:"host" toml:"host"`
	Port     string `json:"port" yaml:"port" toml:"port"`
	LogLevel int    `json:"log_level" yaml:"log_level" toml:"log_level"`
}

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

	cfg := Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
