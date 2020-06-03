package apiv1

import (
	"net/http"

	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/utils/locale"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	msgUserFromJSON = &i18n.Message{ID: "api.user.create_user.json.app_error", Other: "could not decode user json data"}
)

// InitUser inits the user routes
func InitUser(a *API) {
	a.BaseRoutes.Users.Post("/", a.createUser)
	a.BaseRoutes.Users.Post("/login", a.login)
}

func (a *API) createUser(w http.ResponseWriter, r *http.Request) {
	u, e := model.UserFromJSON(r.Body)
	if e != nil {
		respondError(w, model.NewAppErr("createUser", model.ErrInternal, locale.GetUserLocalizer("en"), msgUserFromJSON, http.StatusInternalServerError, nil))
		return
	}

	user, err := a.app.CreateUser(u)
	if err != nil {
		respondError(w, err)
		return
	}

	respondJSON(w, http.StatusCreated, user)
}

func (a *API) login(w http.ResponseWriter, r *http.Request) {
	u, e := model.UserLoginFromJSON(r.Body)
	if e != nil {
		respondError(w, model.NewAppErr("login", model.ErrInternal, locale.GetUserLocalizer("en"), msgUserFromJSON, http.StatusInternalServerError, nil))
		return
	}

	user, err := a.app.Login(u)
	if err != nil {
		respondError(w, err)
		return
	}
	respondJSON(w, http.StatusOK, user)
}
