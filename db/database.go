package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "cotacao.db")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
