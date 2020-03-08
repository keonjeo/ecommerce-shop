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

func (s PGUserStore) Save(*model.User) (*model.User, error) {
	return &model.User{}, nil
}

func (s PGUserStore) Get(id int) (*model.User, error) {
	return &model.User{}, nil
}

func (s PGUserStore) GetAll() ([]*model.User, error) {
	return []*model.User{}, nil
}

func (s PGUserStore) FindByEmail(email string) (*model.User, error) {
	return &model.User{}, nil
}

func (s PGUserStore) Update(id int, u *model.User) (*model.User, error) {
	return &model.User{}, nil
}

func (s PGUserStore) Delete(id int) (*model.User, error) {
	return &model.User{}, nil
}
