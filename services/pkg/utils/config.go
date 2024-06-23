package utils

import (
	"log"
	"net"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// ServerConfig Конфиг сервера
type ServerConfig struct {
	Host    string `yaml:"host" env-required:"true"`
	Port    string `yaml:"port" env-required:"true"`
	Timeout string `yaml:"timeout" env-default:"5"`
}

// ServiceConfig Конфиг Микросервиса
type ServiceConfig struct {
	Server ServerConfig `yaml:"server"`
}

// GlobalConfig Общий глобальный конфиг для всех микросервисов
type GlobalConfig struct {
	Service struct {
		Auth ServiceConfig `yaml:"auth"`
		User ServiceConfig `yaml:"user"`
	}
}

// LoadGlobalConfig Функция загрузки глобального конфига для микросервиса
func LoadGlobalConfig(cfgPath string) *GlobalConfig {
	_, err := os.Stat(cfgPath)
	if err != nil {
		log.Fatalf("Config file doesnt exist. Provided path: %s", cfgPath)
	}

	var cfg GlobalConfig

	if err = cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	return &cfg
}

func (cfg *ServerConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}
