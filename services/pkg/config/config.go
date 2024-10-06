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
	DB  DBConfig
}

// AppConfig Конфиг сервиса.
type AppConfig struct {
	Host    string `env:"APP_HOST"    env-default:"localhost"`
	Port    string `env:"APP_PORT"    env-default:"5431"`
	Timeout string `env:"APP_TIMEOUT" env-default:"5"`
}

// DBConfig Конфиг Базы данных.
type DBConfig struct {
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT"       env-default:"5440"`
	Timeout  string `env:"DB_TIMEOUT"    env-default:"5"`
	User     string `env:"DB_USER"       env-default:"postgres"`
	Password string `env:"DB_PASSWORD"   env-default:"postgres"`
	DBName   string `env:"DB_NAME"       env-default:"postgres"`
}

// ParseConfigFiles gets cfg that implements Config and parses cfg files to extract config fields.
func ParseConfigFiles(cfg interface{}, filePaths ...string) error {
	for i := range filePaths {
		err := cleanenv.ReadConfig(filePaths[i], cfg)
		if err != nil {
			log.Printf("Error reading configuration from file: %v", filePaths[i])

			return err
		}
	}

	return nil
}

func (cfg *Config) GetDataBaseURL() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName)
}

func (cfg *Config) GetMongoDataBaseURL() string {
	creds := fmt.Sprintf("%s:%s", cfg.DB.User, cfg.DB.Password)

	return fmt.Sprintf("mongodb://%s@%s", creds, net.JoinHostPort(cfg.DB.Host, cfg.DB.Port))
}

func (cfg *Config) GetAppAddress() string {
	return net.JoinHostPort(cfg.App.Host, cfg.App.Port)
}
