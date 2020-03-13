package postgres

import (
	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/store"
)

// PGProductStore ...
type PGProductStore struct {
	PGStore
	store.ProductStore
}

// Save ...
func (s PGProductStore) Save(*model.Product) (*model.Product, error) {
	return &model.Product{}, nil
}

// Get ...
func (s PGProductStore) Get(id int) (*model.Product, error) {
	return &model.Product{}, nil
}

// GetAll ...
func (s PGProductStore) GetAll() ([]*model.Product, error) {
	return []*model.Product{}, nil
}

// Update ...
func (s PGProductStore) Update(id int, u *model.Product) (*model.Product, error) {
	return &model.Product{}, nil
}

// Delete ...
func (s PGProductStore) Delete(id int) (*model.Product, error) {
	return &model.Product{}, nil
}
