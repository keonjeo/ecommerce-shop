package store

import (
	"github.com/dankobgd/ecommerce-shop/model"
)

// Store represents all stores
type Store interface {
	User() UserStore
	Product() ProductStore
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

// ProductStore represents the storage for the product model
type ProductStore interface {
	Save(*model.Product) (*model.Product, *model.AppErr)
	Get(id int) (*model.Product, *model.AppErr)
	GetAll() ([]*model.Product, *model.AppErr)
	Update(id int, u *model.Product) (*model.Product, *model.AppErr)
	Delete(id int) (*model.Product, *model.AppErr)
}
