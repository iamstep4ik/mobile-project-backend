package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Server struct {
		Host string `envconfig:"SERVER_HOST"`
		Port string `envconfig:"SERVER_PORT"`
	}
	Database struct {
		Username string `envconfig:"DB_USER"`
		Password string `envconfig:"DB_PASSWORD"`
		Host     string `envconfig:"DB_HOST"`
		Port     string `envconfig:"DB_PORT"`
		Name     string `envconfig:"DB_NAME"`
		SSLMode  string `envconfig:"DB_SSLMODE"`
	}
}

func NewConfig() *Config {
	return &Config{}
}

func LoadConfig(cfg *Config) error {
	return envconfig.Process("", cfg)
}
