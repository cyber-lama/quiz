package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port     string
	LogLevel string
	TTL      string
	DBUrl    string
}

func Init() *Config {
	return &Config{
		Port:     fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		LogLevel: os.Getenv("LOG_LEVEL"),
		TTL:      os.Getenv("TTL"),
		DBUrl: fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME")),
	}
}
