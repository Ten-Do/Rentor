package config

import (
	"errors"
	"os"
	"time"

	"github.com/spf13/viper"
)

type HTTPServer struct {
	Host                string        `yaml:"host"`
	Port                int           `yaml:"port"`
	TimeoutSeconds      time.Duration `yaml:"timeout_seconds"`
	IddleTimeoutSeconds time.Duration `yaml:"iddle_timeout_seconds"`
}

type Config struct {
	Env         string     `yaml:"env"`
	StoragePath string     `yaml:"storage_path"`
	HTTPServer  HTTPServer `yaml:"http_server"`
}

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
