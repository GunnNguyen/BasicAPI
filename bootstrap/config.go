package bootstrap

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Jwt struct {
		TokenTimeLife int    `envconfig:"TOKENTIMELIFE"`
		SecretKey     string `envconfig:"SECRETKEY"`
	}
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	cfg := new(Config)
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("load config error: %v", err)
	}
	return cfg, nil
}
