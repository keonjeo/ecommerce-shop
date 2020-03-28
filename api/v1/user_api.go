package apiv1

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func users(a *API) http.Handler {
	r := chi.NewRouter()

	r.Get("/", a.getUsers)
	r.Get("/{user_id:[A-Za-z0-9]+}", a.getUser)
	r.Post("/", a.postUsers)
	r.Get("/test", a.test)

	return r
}

func (a *API) getUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("all users"))
}

func (a *API) postUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post users"))
}

func (a *API) getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	w.Write([]byte("user: " + userID))
}

func (a *API) test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user from db: " + a.app.Test() + ", cfg port: " + strconv.Itoa(a.app.Cfg().Port)))
}
