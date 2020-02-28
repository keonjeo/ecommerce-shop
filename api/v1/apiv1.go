package apiv1

import (
	"fmt"

	"github.com/go-chi/chi"
)

// Init inits the API
func Init() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/api/v1/", subrouter())
	r.Mount("/api/v1/users", subrouter())
	r.Mount("/api/v1/users/{user_id:[A-Za-z0-9]+}", subrouter())

	return r
}

func subrouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Path))
	})
	return r
}
