package config

import (
	"log"
	"net"

	commonConfig "github.com/Insid1/with-auth/pkg/config"
)

type Config struct {
	commonConfig.Config

	AuthServiceHost string `env:"AUTH_SERVICE_HOST"     env-default:"127.0.0.1"`
	AuthServicePort string `env:"AUTH_SERVICE_PORT"     env-default:"5432"`
}

func MustLoad() *Config {

	configs := []string{"db/.env", ".env"}

	var cfg Config

	err := commonConfig.ParseConfigFiles(&cfg, configs...)
	if err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	return &cfg
}

func (c *Config) GetAuthServiceAddress() string {
	return net.JoinHostPort(c.AuthServiceHost, c.AuthServicePort)
}
