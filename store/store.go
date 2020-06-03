package store

import (
	"github.com/dankobgd/ecommerce-shop/model"
)

// Store represents all stores
type Store interface {
	User() UserStore
}

// UserStore represents the storage for the user model
type UserStore interface {
	Save(*model.User) (*model.User, *model.AppErr)
	Get(id int) (*model.User, *model.AppErr)
	GetAll() ([]*model.User, *model.AppErr)
	GetByEmail(email string) (*model.User, *model.AppErr)
	Update(id int, u *model.User) (*model.User, *model.AppErr)
	Delete(id int) (*model.User, *model.AppErr)
}
