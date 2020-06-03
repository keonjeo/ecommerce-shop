package app

import (
	"net/http"
	"os"

	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/utils/locale"
	"github.com/dgrijalva/jwt-go/v4"
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

// CreateUserAccessToken creates the access token for the user
func CreateUserAccessToken(authD interface{}) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
