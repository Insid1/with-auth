package config

import "github.com/joho/godotenv"

type EnvConfig struct {
}

func NewEnv() (*EnvConfig, error) {
	c := EnvConfig{}

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &c, nil
}
