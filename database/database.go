package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "quotation.db")

	if err != nil {
		return nil, err
	}

	err = database.Ping()

	if err != nil {
		return nil, err
	}

	return database, nil
}
