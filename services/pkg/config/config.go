package config

import (
	"fmt"
	"log"
	"net"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `env:"ENV" env-default:"local"`

	App AppConfig
	Db  DBConfig
}

// AppConfig Конфиг сервиса
type AppConfig struct {
	Host    string `env:"APP_HOST" env-default:"localhost"`
	Port    string `env:"APP_PORT" env-default:"5431"`
	Timeout string `env:"APP_TIMEOUT" env-default:"5"`
}

// DBConfig Конфиг Базы данных
type DBConfig struct {
	Host     string `env:"POSTGRES_HOST"     env-default:"localhost"`
	Port     string `env:"DB_PORT"     env-default:"5441"`
	Timeout  string `env:"DB_TIMEOUT" env-default:"5"`
	User     string `env:"POSTGRES_USER"     env-default:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"postgres"`
	DBName   string `env:"POSTGRES_DB"       env-default:"postgres"`
}

func ParseConfigFiles[C Config](filePaths ...string) (*C, error) {
	var cfg C

	for i := 0; i < len(filePaths); i++ {
		err := cleanenv.ReadConfig(filePaths[i], &cfg)
		if err != nil {
			log.Printf("Error reading configuration from file: %v", filePaths[i])
			return nil, err
		}
	}

	return &cfg, nil
}

func (cfg *Config) GetDataBaseURL() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Db.Host, cfg.Db.Port, cfg.Db.User, cfg.Db.Password, cfg.Db.DBName)
}

func (cfg *Config) AppAddress() string {
	return net.JoinHostPort(cfg.App.Host, cfg.App.Port)
}
