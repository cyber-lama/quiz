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
	var err error
	var app *server.App

	c := config.Init()

	app, err = server.NewApp(c)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(c.Port)
	if err != nil {
		log.Fatal(err)
	}
}
