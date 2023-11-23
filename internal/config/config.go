package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type DbConfig struct {
	Port     string `yaml:"port" env:"PORT" env-default:"5432"`
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	DbName     string `yaml:"db-name" env:"NAME" env-default:"postgres"`
	User     string `yaml:"user" env-default:"user"`
	Password string `yaml:"password" env:"PASSWORD"`
}

type AppConfig struct {
	Db      DbConfig
	AppPort string `yaml:"app-port" env:"APP_PORT" env-default:"1234"`
}

var config AppConfig
var once sync.Once

func GetConfig() AppConfig {
	once.Do(func() {
		err := cleanenv.ReadConfig("internal/config/app-config.yml", &config)
		if err != nil {
			panic(err)
		}
	})
	return config
}
