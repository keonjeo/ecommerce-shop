package apiv1

import (
	"github.com/dankobgd/ecommerce-shop/app"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// API ...
type API struct {
	app *app.App
}

// Init inits the API
func Init(a *app.App, r *chi.Mux) {
	api := &API{
		app: a,
	}

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/api/v1/users", users(api))
}
