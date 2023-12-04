package database

import "database/sql"

type database struct {
	db *sql.DB
}

func NewDatabase(db *sql.DB) database {
	return database{
		db: db,
	}
}
