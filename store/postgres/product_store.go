package postgres

import (
	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/store"
)

// PGProductStore ...
type PGProductStore struct {
	PGStore
}

func newPGProductStore() store.ProductStore {
	return &PGProductStore{}
}

func (s PGProductStore) Save(*model.Product) (*model.Product, error) {
	return &model.Product{}, nil
}

func (s PGProductStore) Get(id int) (*model.Product, error) {
	return &model.Product{}, nil
}

func (s PGProductStore) GetAll() ([]*model.Product, error) {
	return []*model.Product{}, nil
}

func (s PGProductStore) Update(id int, u *model.Product) (*model.Product, error) {
	return &model.Product{}, nil
}

func (s PGProductStore) Delete(id int) (*model.Product, error) {
	return &model.Product{}, nil
}
