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

func (s PGStore) Connect() {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.Query("SELECT 1;")
}

func (s PGStore) Close() {}

func (s PGStore) New() {

}
