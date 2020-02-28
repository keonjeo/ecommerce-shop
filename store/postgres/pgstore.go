package postgres

import (
	"github.com/jmoiron/sqlx"

	"github.com/dankobgd/ecommerce-shop/store"
)

// PGStore ...
type PGStore struct {
	db *sqlx.DB
	store.Store
}

// SQLStore ...
type SQLStore struct {
	user    store.UserStore
	product store.ProductStore
}
