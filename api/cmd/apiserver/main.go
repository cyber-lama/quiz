package main

import (
	"api/internal/exceptions"
	app "api/internal/kernel"
	"errors"
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
	server, err := app.Init()
	if err != nil {
		if errors.As(err, &exceptions.ConnectionDBErr{}) {
			log.Fatal(err.Error())
		} else {
			log.Fatalf("Ошибка инициализации приложения %s", err)
		}
	}

	if err = server.Run(); err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}
