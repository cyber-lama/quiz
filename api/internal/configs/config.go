package configs

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

func Run() *Config {
	return &Config{
		Port:     fmt.Sprintf(":%s", os.Getenv("API_PORT")),
		LogLevel: os.Getenv("LOG_LEVEL"),
		TTL:      os.Getenv("TTL"),
		DBUrl: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")),
	}
}
