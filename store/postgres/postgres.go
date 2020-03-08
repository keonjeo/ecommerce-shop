package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/dankobgd/ecommerce-shop/store"
)

// PGStore ...
type PGStore struct {
	db *sqlx.DB
	store.Store
}

// New ...
func (s PGStore) New() (*PGStore, error) {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &PGStore{db: db}, nil
}
