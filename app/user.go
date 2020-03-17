package app

import (
	"github.com/dankobgd/ecommerce-shop/model"
)

// CreateUser ...
func (a *App) CreateUser(user *model.User) (*model.User, error) {
	// a.Srv().Store.User().Save(id)
	// user, err := a.Srv().Store.User().GetByEmail(email)
	return nil, nil
}

// Test ...
func (a *App) Test() string {
	return a.Srv().Store.User().Test()
}
