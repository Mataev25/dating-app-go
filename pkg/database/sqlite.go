package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteDB(filepath string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	db.Exec("PRAGMA foreign_keys = ON")

	return db, nil
}
