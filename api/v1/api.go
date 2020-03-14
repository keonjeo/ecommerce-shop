package apiv1

import (
	"github.com/go-chi/chi"
)

// Init inits the API
func Init(r *chi.Mux) {
	r.Mount("/api/v1", users())
	r.Mount("/api/v1/users", users())
	r.Mount("/api/v1/users/{user_id:[A-Za-z0-9]+}", userByID())
}
