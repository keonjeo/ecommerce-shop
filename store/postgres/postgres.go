package postgres

import (
	"log"

	"github.com/dankobgd/ecommerce-shop/store"
	_ "github.com/jackc/pgx/stdlib" // pg driver
	"github.com/jmoiron/sqlx"
)

// PgStore ...
type PgStore struct {
	db *sqlx.DB
	store.Store
}

// Connect establishes connection to postgres db
func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", "host=localhost user=test password=test dbname=ecommerce sslmode=disable")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return db, nil
}

// New initializes postgres based store
func (s PgStore) New(db *sqlx.DB) *PgStore {
	return &PgStore{db: db}
}
