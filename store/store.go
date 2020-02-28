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
	Save(*model.User) (*model.User, error)
	Get(id int) (*model.User, error)
	GetAll() ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(id int, u *model.User) (*model.User, error)
	Delete(id int) (*model.User, error)
}

// ProductStore represents the storage for the product model
type ProductStore interface {
	Save(*model.Product) (*model.Product, error)
	Get(id int) (*model.Product, error)
	GetAll() ([]*model.Product, error)
	Update(id int, u *model.Product) (*model.Product, error)
	Delete(id int) (*model.Product, error)
}
