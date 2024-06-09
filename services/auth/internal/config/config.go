package config

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env    string `env:"ENV" env-default:"local"`
	JWTKey string `env:"JWT_SECRET_KEY" env-required:"true"`
	Db     DBConfig
	Grpc   GRPCConfig
}

type DBConfig struct {
	Host     string `env:"POSTGRES_HOST"     env-default:"localhost"`
	Port     string `env:"POSTGRES_PORT"     env-default:"5433"`
	User     string `env:"POSTGRES_USER"     env-required:"true"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
	DBName   string `env:"POSTGRES_DB"       env-required:"true"`
}

type GRPCConfig struct {
	Host    string `env:"GRPC_HOST"    env-default:"localhost"`
	Port    string `env:"GRPC_PORT"    env-default:"5433"`
	Timeout string `env:"GRPC_TIMEOUT" env-default:"5s"`
}

func MustLoad() *Config {

	cfgPath := ".env"

	_, err := os.Stat(cfgPath)
	if err != nil {
		log.Fatalf("Config file doesnt exist. Provided path: %s", cfgPath)
	}

	var cfg Config

	if err = cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		log.Fatalf("Error reading config: %s", err)
	}
	return &cfg
}

func (cfg *GRPCConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func (cfg *DBConfig) DataSourceName() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
}
