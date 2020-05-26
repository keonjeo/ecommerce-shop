package apiv1

import (
	"net/http"

	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/utils/locale"
	"github.com/go-chi/chi"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func users(a *API) http.Handler {
	r := chi.NewRouter()

	r.Post("/", a.createUser)
	r.Post("/login", a.login)
	// r.Get("/{user_id:[A-Za-z0-9]+}", a.getUser)

	return r
}

func (a *API) createUser(w http.ResponseWriter, r *http.Request) {
	user, err := model.UserFromJSON(r.Body)
	if err != nil {
		respondError(w, model.NewAppError("createUser", locale.GetUserLocalizer("en"), &i18n.Message{ID: "api.user.create_user.json.app_error"}, err.Error(), http.StatusInternalServerError))
		return
	}

	createdUser, err2 := a.app.CreateUser(user)
	if err2 != nil {
		respondError(w, err2)
		return
	}

	respondJSON(w, http.StatusCreated, createdUser)
}

func (a *API) login(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{"login": "x"})
}
