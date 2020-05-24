package app

import (
	"github.com/dankobgd/ecommerce-shop/model"
)

// CreateUser creates the new user in the system
func (a *App) CreateUser(user *model.User) (*model.User, *model.AppError) {
	user.PreSave()
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user, err := a.Srv().Store.User().Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
