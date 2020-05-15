package app

import (
	"github.com/dankobgd/ecommerce-shop/model"
)

func (a *App) GetUsers() ([]*model.User, error) {
	a.log.Debug("Debug")
	a.log.Info("Info")
	a.log.Warn("Warning")
	a.log.Error("Error")

	return nil, nil
}

// CreateUser ...
func (a *App) CreateUser(user *model.User) (*model.User, error) {
	// a.Srv().Store.User().Save(id)
	// user, err := a.Srv().Store.User().GetByEmail(email)
	return nil, nil
}
