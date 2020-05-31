package app

import (
	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/zlog"
)

// CreateUser creates the new user in the system
func (a *App) CreateUser(user *model.User) (*model.User, *model.AppErr) {
	rawpw := user.Password
	user.PreSave()
	if err := user.Validate(); err != nil {
		return nil, err

	}

	if err := model.IsValidPasswordCriteria(rawpw, &a.Cfg().PasswordSettings); err != nil {
		return nil, err
	}

	user, err := a.Srv().Store.User().Save(user)
	if err != nil {
		a.log.Error(err.Error(), zlog.Err(err))
		return nil, err
	}

	user.Sanitize(map[string]bool{})
	return user, nil
}
