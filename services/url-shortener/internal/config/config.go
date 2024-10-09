package config

import (
	"log"
	"net"

	commonConfig "github.com/Insid1/with-auth/pkg/config"
)

type AppConfig struct {
	AuthServiceHost    string `env:"AUTH_SERVICE_HOST"    env-default:"127.0.0.1"`
	AuthServicePort    string `env:"AUTH_SERVICE_PORT"    env-default:"5432"`
	ShortenerURLPrefix string `env:"SHORTENER_URL_PREFIX" env-default:"https://provide.your.url"`
}

type Config struct {
	commonConfig.Config
	AppConfig
}

func MustLoad() *Config {
	cfgPaths := []string{".env", "db/.env"}

	var cfg Config

	err := commonConfig.ParseConfigFiles(&cfg, cfgPaths...)
	if err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	return &cfg
}

func (cfg *Config) GetAuthServiceAddress() string {
	return net.JoinHostPort(cfg.AppConfig.AuthServiceHost, cfg.AppConfig.AuthServicePort)
}
