package db

import (
	"database/sql"
	"log"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./cotacao.db")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
