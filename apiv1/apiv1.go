package apiv1

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Init inits the API
func Init(r chi.Router) {
	r.Mount("/api/v1/", subrouter())
	r.Mount("/api/v1/users", subrouter())
	r.Mount("/api/v1/users/{user_id:[A-Za-z0-9]+}", subrouter())
}

func subrouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Path))
	})
	return r
}
