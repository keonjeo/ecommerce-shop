package app

import (
	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/zlog"
)

// CreateUser creates the new user in the system
func (a *App) CreateUser(u *model.User) (*model.User, *model.AppErr) {
	rawpw := u.Password
	u.PreSave()
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := a.IsValidPassword(rawpw); err != nil {
		return nil, err
	}

	user, err := a.Srv().Store.User().Save(u)
	if err != nil {
		a.log.Error(err.Error(), zlog.Err(err))
		return nil, err
	}

	user.Sanitize(map[string]bool{})
	return user, nil
}

// Login handles the user login
func (a *App) Login(u *model.UserLogin) (*model.User, *model.AppErr) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	user, err := a.Srv().Store.User().GetByEmail(u.Email)
	if err != nil {
		a.log.Error(err.Error(), zlog.Err(err))
		return nil, err
	}
	if err := a.CheckUserPassword(user, u.Password); err != nil {
		return nil, err
	}

	user.Sanitize(map[string]bool{})
	return user, nil
}
