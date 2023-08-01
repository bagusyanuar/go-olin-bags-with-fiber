package config

import (
	"errors"
	"fmt"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Config struct {
	Host       string `env:"SERVER_HOST,default=localhost"`
	Port       string `env:"SERVER_PORT"`
	AppName    string `env:"APP_NAME"`
	AppVersion string `env:"APP_VERSION"`
	AppMode    string `env:"APP_MODE,default=dev"`
}

func NewConfig(env string) (*Config, error) {
	_ = godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		message := fmt.Sprintf("error load %s file", env)
		return nil, errors.New(message)
	}
	return &config, nil
}
