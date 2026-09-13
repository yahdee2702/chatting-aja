package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Env  string
	Port string
}

func Load() *Config {
	godotenv.Load()
	return &Config{
		App: AppConfig{
			Env:  os.Getenv("APP_ENV"),
			Port: os.Getenv("APP_PORT"),
		},
	}
}
