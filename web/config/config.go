package config

import "github.com/caarlos0/env"

type Config struct {
	LOGIN_URL        string `env:"LOGIN_URL"`
	VERIFY_TOKEN_URL string `env:"VERIFY_TOKEN_URL"`
	SERVER_URL       string `env:"SERVER_URL"`
	HTTP_PORT        int    `env:"HTTP_PORT" envDefault:"8080"`
	APP_TITLE        string `env:"APP_TITLE" envDefault:"Web Client"`
}

func Parse() *Config {
	cfg := new(Config)
	env.Parse(cfg)
	return cfg
}
