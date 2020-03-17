package postgres

import (
	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/store"
)

// PgUserStore ...
type PgUserStore struct {
	PgStore
}

func newPgUserStore(pgst *PgStore) store.UserStore {
	return &PgUserStore{*pgst}
}

// Test ...
func (s PgUserStore) Test() string {
	var str string
	_ = s.db.Get(&str, "select * from user;")
	return str
}

// Save ...
func (s PgUserStore) Save(*model.User) (*model.User, error) {
	return &model.User{}, nil
}

// Get ...
func (s PgUserStore) Get(id int) (*model.User, error) {
	return &model.User{}, nil
}

// GetAll ...
func (s PgUserStore) GetAll() ([]*model.User, error) {
	return []*model.User{}, nil
}

// GetByEmail ...
func (s PgUserStore) GetByEmail(email string) (*model.User, error) {
	return &model.User{}, nil
}

// Update ...
func (s PgUserStore) Update(id int, u *model.User) (*model.User, error) {
	return &model.User{}, nil
}

// Delete ...
func (s PgUserStore) Delete(id int) (*model.User, error) {
	return &model.User{}, nil
}
