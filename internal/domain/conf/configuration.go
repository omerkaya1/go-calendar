package conf

import (
	"github.com/omerkaya1/go-calendar/internal/domain/errors"
	"github.com/spf13/viper"
	"path"
)

// MEMO: move to models?
type Config struct {
	Host     string `json:"host" yaml:"host" toml:"host"`
	Port     string `json:"port" yaml:"port" toml:"port"`
	LogLevel int    `json:"log_level" yaml:"log_level" toml:"log_level"`
	DB       DBConf `json:"db" yaml:"db" toml:"db"`
}

type DBConf struct {
	Name    string `json:"dbname" yaml:"dbname" toml:"dbname"`
	User    string `json:"user" yaml:"user" toml:"user"`
	SSLMode string `json:"sslmode" yaml:"sslmode" toml:"sslmode"`
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

func CfgFromFile(cfgPath string) (*Config, error) {
	cfg, err := InitConfig(cfgPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func CfgFromCmdParams(logLevel int, host, port, dbName, dbUser, sslMode string) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		LogLevel: logLevel,
		DB: DBConf{
			Name:    dbName,
			User:    dbUser,
			SSLMode: sslMode,
		},
	}
}
