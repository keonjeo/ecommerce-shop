package apiv1

import (
	"net/http"

	"github.com/go-chi/chi"
)

func users() http.Handler {
	r := chi.NewRouter()
	r.Get("/", getUsers)
	r.Post("/", postUsers)
	return r
}

func userByID() http.Handler {
	r := chi.NewRouter()
	r.Get("/", getUser)
	r.Get("/invoices", getUserInvoices)
	return r
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func postUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func getUserInvoices(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}
