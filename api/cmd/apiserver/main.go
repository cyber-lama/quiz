package main

import (
	"api/internal/exceptions"
	"api/internal/server"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
func Print[T string](s ...T) {
	for _, v := range s {
		fmt.Print(v)
	}
}
func main() {
	s, err := server.Init()
	if err != nil {
		if errors.As(err, &exceptions.ConnectionDBErr{}) {
			log.Fatal(err.Error())
		} else {
			log.Fatalf("Application initialization error %s", err)
		}
	}

	s.Run()
}
