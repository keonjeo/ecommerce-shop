package apiv1

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/utils/locale"
	"github.com/go-chi/chi"
)

func users(a *API) http.Handler {
	r := chi.NewRouter()

	r.Get("/", a.test)
	r.Post("/", a.createUser)
	r.Post("/login", a.login)
	// r.Get("/{user_id:[A-Za-z0-9]+}", a.getUser)

	return r
}

func (a *API) createUser(w http.ResponseWriter, r *http.Request) {
	user, err := model.UserFromJSON(r.Body)

	if err != nil {
		respondError(w, model.NewAppError("createUser", "api.user.create_user", nil, err.Error(), http.StatusInternalServerError))
		return
	}

	if err := user.Validate(); err != nil {
		a.app.Log().Error(err.Error())
		// http.Error(w, err.Error(), http.StatusBadRequest)

		respondError(w, model.NewAppError("createUser", "api.user.create_user", nil, err.Error(), http.StatusInternalServerError))
		return
	}

	respondJSON(w, http.StatusCreated, user)
}

func (a *API) login(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{"login": "x"})
}

func (a *API) test(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")
	msg := locale.T("hello", struct{ Person string }{Person: "Bob"})

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	if n > 5 {
		respondError(w, model.NewAppError("test", "api.test.app_error", nil, "", http.StatusInternalServerError))
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"msg": msg, "lang": lang})
	return
}
