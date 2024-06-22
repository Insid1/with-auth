package config

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string `env:"ENV" env-default:"local"`
	JWTKey   string `env:"JWT_SECRET_KEY" env-required:"true"`
	Db       DBConfig
	Grpc     GRPCConfig
	Security SecurityConfig
}

type DBConfig struct {
	Host     string `env:"POSTGRES_HOST"     env-default:"localhost"`
	Port     string `env:"POSTGRES_PORT"     env-default:"5432"`
	User     string `env:"POSTGRES_USER"     env-required:"true"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
	DBName   string `env:"POSTGRES_DB"       env-required:"true"`
}

type GRPCConfig struct {
	Host    string `env:"GRPC_HOST"    env-default:"localhost"`
	Port    string `env:"GRPC_PORT"    env-default:"5433"`
	Timeout string `env:"GRPC_TIMEOUT" env-default:"5s"`
}

type SecurityConfig struct {
	ServerCertPath string `env:"SERVER_CERT_PATH"     env-required:"true"`
	ServerKeyPath  string `env:"SERVER_KEY_PATH"     env-required:"true"`

	ClientCertPath string `env:"CLIENT_CERT_PATH"     env-required:"true"`
	ClientKeyPath  string `env:"CLIENT_KEY_PATH"     env-required:"true"`

	CAClientCertPath string `env:"CA_CLIENT_CERT_PATH"     env-required:"true"`
	CAServerCertPath string `env:"CA_SERVER_CERT_PATH"     env-required:"true"`
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
