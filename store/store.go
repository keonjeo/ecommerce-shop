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
	Save(*model.User) (*model.User, *model.AppError)
	Get(id int) (*model.User, *model.AppError)
	GetAll() ([]*model.User, *model.AppError)
	GetByEmail(email string) (*model.User, *model.AppError)
	Update(id int, u *model.User) (*model.User, *model.AppError)
	Delete(id int) (*model.User, *model.AppError)
}

// ProductStore represents the storage for the product model
type ProductStore interface {
	Save(*model.Product) (*model.Product, *model.AppError)
	Get(id int) (*model.Product, *model.AppError)
	GetAll() ([]*model.Product, *model.AppError)
	Update(id int, u *model.Product) (*model.Product, *model.AppError)
	Delete(id int) (*model.Product, *model.AppError)
}
