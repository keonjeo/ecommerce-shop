package postgres

import (
	"log"

	_ "github.com/jackc/pgx" // pg driver
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
	db, err := sqlx.Connect("postgres", "host=database port=5432 user=postgres password=postgres dbname=ecommerce sslmode=disable")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &PGStore{db: db}, nil
}
