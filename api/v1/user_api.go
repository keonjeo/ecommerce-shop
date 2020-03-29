package apiv1

import (
	"net/http"

	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/go-chi/chi"
)

func users(a *API) http.Handler {
	r := chi.NewRouter()

	r.Get("/", a.getUsers)
	r.Post("/", a.postUsers)
	r.Get("/{user_id:[A-Za-z0-9]+}", a.getUser)

	return r
}

func (a *API) getUsers(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, a.app.Cfg())
}

func (a *API) postUsers(w http.ResponseWriter, r *http.Request) {
	user := model.UserFromJSON(r.Body)

	respondJSON(w, http.StatusCreated, user)
}

func (a *API) getUser(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, chi.URLParam(r, "user_id"))
}
