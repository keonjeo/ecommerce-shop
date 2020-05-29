package postgres

import (
	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/store"
)

// PgProductStore ...
type PgProductStore struct {
	PgStore
}

func newPgProductStore(pgst *PgStore) store.ProductStore {
	return &PgProductStore{*pgst}
}

// Save ...
func (s PgProductStore) Save(*model.Product) (*model.Product, *model.AppErr) {
	return &model.Product{}, nil
}

// Get ...
func (s PgProductStore) Get(id int) (*model.Product, *model.AppErr) {
	return &model.Product{}, nil
}

// GetAll ...
func (s PgProductStore) GetAll() ([]*model.Product, *model.AppErr) {
	return []*model.Product{}, nil
}

// Update ...
func (s PgProductStore) Update(id int, u *model.Product) (*model.Product, *model.AppErr) {
	return &model.Product{}, nil
}

// Delete ...
func (s PgProductStore) Delete(id int) (*model.Product, *model.AppErr) {
	return &model.Product{}, nil
}
