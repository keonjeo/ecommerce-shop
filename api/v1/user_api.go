package apiv1

import (
	"log"
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
	a.app.GetUsers()
	respondError(w, model.NewAppError("GetUsers", "api.user_api.get_users", "ErrGetUsers", "could not get users", map[string]interface{}{"a": "b"}, http.StatusInternalServerError))
	return
}

func (a *API) postUsers(w http.ResponseWriter, r *http.Request) {
	user, err := model.UserFromJSON(r.Body)
	if err != nil {
		respondError(w, model.NewAppError("postUsers", "api.user_api.post_users", "ErrPostUsers", err.Error(), nil, http.StatusInternalServerError))
		return
	}

	if err := user.Validate(); err != nil {
		log.Printf("%v\n", err)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respondJSON(w, http.StatusCreated, user)
}

func (a *API) getUser(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, chi.URLParam(r, "user_id"))
}
