package config

import (
	"errors"
	"os"
	"time"

	"github.com/spf13/viper"
)

// TODO: add more configuration fields as needed

type HTTPServer struct {
	Host               string        `mapstructure:"host" yaml:"host"`
	Port               string        `mapstructure:"port" yaml:"port"`
	TimeoutSeconds     time.Duration `mapstructure:"timeout_seconds" yaml:"timeout_seconds"`
	IdleTimeoutSeconds time.Duration `mapstructure:"idle_timeout_seconds" yaml:"idle_timeout_seconds"`
}

type Auth struct {
	JWTSecret            string        `mapstructure:"jwt_secret" yaml:"jwt_secret" default:""`
	AccessTokenTTL       time.Duration `mapstructure:"access_token_ttl" yaml:"access_token_ttl"`
	RefreshTokenTTL      time.Duration `mapstructure:"refresh_token_ttl" yaml:"refresh_token_ttl"`
	OTPLength            int           `mapstructure:"otp_length" yaml:"otp_length"`
	OTPExpirationMinutes int           `mapstructure:"otp_expiration_minutes" yaml:"otp_expiration_minutes"`
	OTPMaxAttempts       int           `mapstructure:"otp_max_attempts" yaml:"otp_max_attempts"`
}

type SMTP struct {
	SMTPFrom     string `mapstructure:"smtp_from" yaml:"smtp_from" default:""`
	SMTPPassWord string `mapstructure:"smtp_password" yaml:"smtp_password" default:""`
	SMTPHost     string `mapstructure:"smtp_host" yaml:"smtp_host" default:""`
	SMTPPort     string `mapstructure:"smtp_port" yaml:"smtp_port" default:""`
}

type Config struct {
	Env              string     `mapstructure:"env" yaml:"env"`
	StoragePath      string     `mapstructure:"storage_path" yaml:"storage_path"`
	ImageStoragePath string     `mapstructure:"image_storage_path" yaml:"image_storage_path"`
	BaseURL          string     `mapstructure:"base_url" yaml:"base_url"`
	HTTPServer       HTTPServer `mapstructure:"http_server" yaml:"http_server"`
	Auth             Auth       `mapstructure:"auth" yaml:"auth"`
	SMTP             SMTP       `mapstructure:"smtp" yaml:"smtp"`
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

	if config.Auth.JWTSecret == "" {
		config.Auth.JWTSecret = os.Getenv("JWT_SECRET")
		if config.Auth.JWTSecret == "" {
			return nil, errors.New("LoadConfig: JWT_SECRET env variable is not set")
		}
	}

	if config.SMTP.SMTPFrom == "" {
		config.SMTP.SMTPFrom = os.Getenv("SMTP_FROM")
		if config.SMTP.SMTPFrom == "" {
			return nil, errors.New("LoadConfig: SMTP_FROM env variable is not set")
		}
	}

	if config.SMTP.SMTPPassWord == "" {
		config.SMTP.SMTPPassWord = os.Getenv("SMTP_PASSWORD")
		if config.SMTP.SMTPPassWord == "" {
			return nil, errors.New("LoadConfig: SMTP_PASSWORD env variable is not set")
		}
	}

	if config.SMTP.SMTPHost == "" {
		config.SMTP.SMTPHost = os.Getenv("SMTP_HOST")
		if config.SMTP.SMTPHost == "" {
			return nil, errors.New("LoadConfig: SMTP_HOST env variable is not set")
		}
	}

	if config.SMTP.SMTPPort == "" {
		config.SMTP.SMTPPort = os.Getenv("SMTP_PORT")
		if config.SMTP.SMTPPort == "" {
			return nil, errors.New("LoadConfig: SMTP_PORT env variable is not set")
		}
	}

	if os.Getenv("DOCKER") == "true" {
		config.HTTPServer.Host = "0.0.0.0"
	}

	return &config, nil
}
