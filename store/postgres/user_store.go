package postgres

import (
	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/store"
)

// PGUserStore ...
type PGUserStore struct {
	PGStore
}

func newPGUserStore() store.UserStore {
	return &PGUserStore{}
}

// Save ...
func (s PGUserStore) Save(*model.User) (*model.User, error) {
	return &model.User{}, nil
}

// Get ...
func (s PGUserStore) Get(id int) (*model.User, error) {
	return &model.User{}, nil
}

// GetAll ...
func (s PGUserStore) GetAll() ([]*model.User, error) {
	return []*model.User{}, nil
}

// GetByEmail ...
func (s PGUserStore) GetByEmail(email string) (*model.User, error) {
	return &model.User{}, nil
}

// Update ...
func (s PGUserStore) Update(id int, u *model.User) (*model.User, error) {
	return &model.User{}, nil
}

// Delete ...
func (s PGUserStore) Delete(id int) (*model.User, error) {
	return &model.User{}, nil
}
