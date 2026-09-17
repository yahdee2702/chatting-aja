package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig
	DB  DatabaseConfig
}

type AppConfig struct {
	Env    string
	Port   string
	Secret string
}

type DatabaseConfig struct {
	Driver   string
	Scheme   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func Load() *Config {
	godotenv.Load()
	return &Config{
		App: AppConfig{
			Env:    os.Getenv("APP_ENV"),
			Port:   os.Getenv("APP_PORT"),
			Secret: os.Getenv("APP_SECRET"),
		},
		DB: DatabaseConfig{
			Driver:   os.Getenv("DB_DRIVER"),
			Scheme:   os.Getenv("DB_SCHEME"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
	}
}
