package main

import (
	app "api/internal/kernel"
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
	server := app.Init()
	server.Run()
}
