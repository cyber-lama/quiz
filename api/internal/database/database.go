package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	*sqlx.DB
}

func Connect(DBUrl string) *DB {
	fmt.Println(DBUrl)
	db, err := sqlx.Connect("postgres",
		DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &DB{
		db,
	}
}
