package config

import (
	"errors"
	"os"
	"time"

	"github.com/spf13/viper"
)

type HTTPServer struct {
	Host                string        `mapstructure:"host" yaml:"host"`
	Port                int           `mapstructure:"port" yaml:"port"`
	TimeoutSeconds      time.Duration `mapstructure:"timeout_seconds" yaml:"timeout_seconds"`
	IddleTimeoutSeconds time.Duration `mapstructure:"iddle_timeout_seconds" yaml:"iddle_timeout_seconds"`
}

type Config struct {
	Env         string     `mapstructure:"env" yaml:"env"`
	StoragePath string     `mapstructure:"storage_path" yaml:"storage_path"`
	HTTPServer  HTTPServer `mapstructure:"http_server" yaml:"http_server"`
}

// LoadConfig loads configuration from a YAML file specified by CONFIG_PATH env variable
func LoadConfig() (*Config, error) {
	config_path := os.Getenv("CONFIG_PATH")
	if config_path == "" {
		return nil, errors.New("LoadConfig: CONFIG_PATH env variable is not set")
	}

	viper.SetConfigFile(config_path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New("LoadConfig: error reading config file: " + err.Error())
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.New("LoadConfig: error unmarshalling config: " + err.Error())
	}
	return &config, nil
}
