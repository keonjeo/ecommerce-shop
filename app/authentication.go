package app

import (
	"net/http"

	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/utils/locale"
)

// IsValidPassword checks if user password is valid
func (a *App) IsValidPassword(password string) *model.AppErr {
	return model.IsValidPasswordCriteria(password, &a.Cfg().PasswordSettings)
}

// CheckUserPassword checks if password matches the hashed version
func (a *App) CheckUserPassword(user *model.User, password string) *model.AppErr {
	if !model.ComparePassword(user.Password, password) {
		return model.NewAppErr("App.ComparePassword", model.ErrConflict, locale.GetUserLocalizer("en"), model.MsgComparePwd, http.StatusBadRequest, nil)
	}
	return nil
}
