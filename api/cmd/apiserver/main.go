package main

import (
	"api/internal/config"
	"api/internal/server"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	c := config.Init()
	_, err := server.NewApp(c)
	if err != nil {
		log.Fatal(err)
	}
}
