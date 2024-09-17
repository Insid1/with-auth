package config

import (
	"log"

	commonConfig "github.com/Insid1/go-auth-user/pkg/config"
)

type Config struct {
	*commonConfig.Config
}

func MustLoad() *Config {

	configs := []string{"db/.env", ".env"}

	cfg, err := commonConfig.ParseConfigFiles(configs...)
	if err != nil {
		log.Fatalf("Error reading config: %s", err)

	}

	return &Config{
		Config: cfg,
	}
}
