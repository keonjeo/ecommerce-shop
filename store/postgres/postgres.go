package postgres

import (
	"log"

	"github.com/dankobgd/ecommerce-shop/store"
	_ "github.com/jackc/pgx/stdlib" // pg driver
	"github.com/jmoiron/sqlx"
)

// PgStore has the pg db driver
type PgStore struct {
	db     *sqlx.DB
	stores providedStores
}

type providedStores struct {
	user    store.UserStore
	product store.ProductStore
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

// NewStore initializes postgres based store
func NewStore(db *sqlx.DB) *PgStore {
	pgst := &PgStore{db: db}

	pgst.stores.user = newPgUserStore(pgst)
	pgst.stores.product = newPgProductStore(pgst)

	return pgst
}

func (s *PgStore) User() store.UserStore {
	return s.stores.user
}

func (s *PgStore) Product() store.ProductStore {
	return s.stores.product
}
