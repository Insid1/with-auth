package config

import (
	"log"
	"net"

	commonConfig "github.com/Insid1/go-auth-user/pkg/config"
)

type AppConfig struct {
	commonConfig.AppConfig
	JwtSecretKey    string `env:"JWT_SECRET_KEY"     env-default:"some-test-secret"`
	UserServiceHost string `env:"USER_SERVICE_HOST"     env-default:"127.0.0.1"`
	UserServicePort string `env:"USER_SERVICE_PORT"     env-default:"5431"`
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

func (cfg *Config) GetUserServiceAddress() string {
	return net.JoinHostPort(cfg.AppConfig.UserServiceHost, cfg.AppConfig.UserServicePort)
}
