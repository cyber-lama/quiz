package database

import (
	"api/internal/exceptions"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Database struct {
	*sqlx.DB
}

func Connect(DBUrl string) (*Database, error) {
	// цикл подключения к БД. Пытаемся 3 раза, если не удалось подсоединиться с первого раза.
	var err error
	var db *sqlx.DB
	tries := 0
	for {
		if tries > 2 {
			err = exceptions.NewConnectionErr(err)
			return nil, err
		}
		db, err = sqlx.Connect("postgres", DBUrl)
		if err != nil {
			// error of connect to db, wait 1s and try to reconnect
			time.Sleep(1 * time.Second)
			tries++
			continue
		} else {
			return &Database{db}, nil
		}
	}
}

func (d *Database) CloseDBConnect() {
	fmt.Println("close")
	err := d.Close()
	if err != nil {
		log.Fatal(err)
	}
}
