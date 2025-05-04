package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)


func NewDatabase(databaseURL string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", databaseURL)
	if err != nil {
		log.Fatalf("cannot start postgres database: %v", err)
	}

	return db

}
