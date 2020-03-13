package postgres

import (
	"log"

	"github.com/dankobgd/ecommerce-shop/store"
	_ "github.com/jackc/pgx" // pg driver
	"github.com/jmoiron/sqlx"
)

// PGStore ...
type PGStore struct {
	db *sqlx.DB
	store.Store
}

// Connect establishes connection to postgres db
func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "host=database port=5432 user=test password=test dbname=ecommerce sslmode=disable")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return db, nil
}

// New initializes postgres based store
func (s PGStore) New(db *sqlx.DB) (*PGStore, error) {
	return &PGStore{db: db}, nil
}
