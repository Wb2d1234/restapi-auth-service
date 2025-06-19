package config

import (
	"os"
)

type Config struct {
	DB_DSN     string
	JWT_SECRET string
	PORT       string
}

func Load() (*Config, error) {
	return &Config{
		DB_DSN:     os.Getenv("DB_DSN"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
		PORT:       os.Getenv("PORT"),
	}, nil
}
