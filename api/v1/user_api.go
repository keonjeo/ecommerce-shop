package apiv1

import (
	"net/http"

	"github.com/go-chi/chi"
)

func users(api *API) http.Handler {
	r := chi.NewRouter()

	r.Get("/", api.getUsers)
	r.Get("/{user_id:[A-Za-z0-9]+}", api.getUser)
	r.Post("/", api.postUsers)
	r.Get("/test", api.test)

	return r
}

func (api *API) getUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("all users"))
}

func (api *API) postUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post users"))
}

func (api *API) getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	w.Write([]byte("user: " + userID))
}

func (api *API) test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from db: " + api.app.Test()))
}
