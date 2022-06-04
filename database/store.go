package database

import (
	"database/sql"
)

type Store struct {
	Db      *sql.DB
	Querier Querier
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Db: db,
	}
}