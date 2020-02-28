package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dankobgd/ecommerce-shop/apiv1"
	"github.com/dankobgd/ecommerce-shop/store"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

// Server ...
type Server struct {
	Store  store.Store
	Server *http.Server
	Router *chi.Mux
	// Log *log.Logger
	// RateLimiter *RateLimiter
	// jobs
	// other cfg
}

// NewServer ...
func NewServer() (*Server, error) {
	router := chi.NewRouter()
	apiv1.Init(router)

	s := &Server{
		Router: router,
	}

	// s.Log = log.NewLogger()

	return s, nil
}

// Start runs the HTTP server
func (s *Server) Start(ctx context.Context) (err error) {
	var handler http.Handler = s.Router

	// TODO: read from cfg
	corsWrapper := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Content-Length", "Cache-Control", "Content-Language", "Content-Type", "Expires", "Last-Modified", "Pragma", "Authorization"},
		MaxAge:           86400,
		AllowCredentials: true,
		Debug:            false,
	})

	handler = corsWrapper.Handler(handler)

	// if ratelimit set it here...

	listenAddr := ":3001"

	s.Server = &http.Server{
		// ErrorLog: logger
		Handler:           handler,
		Addr:              listenAddr,
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      7 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	// autocert manager later...
	// if UseLetsEncrypt...

	go func() {
		if err := s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not listen on %s: %v\n", listenAddr, err)
		}
	}()

	log.Printf("server started")

	<-ctx.Done()

	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = s.Server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server shutdown failed: %+s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}

// Stop stops the HTTP server
func (s *Server) Stop() {

}

// package app

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/dankobgd/ecommerce-shop/apiv1"
// 	"github.com/dankobgd/ecommerce-shop/store"
// 	"github.com/go-chi/chi"
// 	"github.com/rs/cors"
// )

// // Server ...
// type Server struct {
// 	Store  store.Store
// 	Server *http.Server
// 	Router *chi.Mux
// 	// RateLimiter *RateLimiter
// }

// // NewServer ...
// func NewServer() (*Server, error) {
// 	// r := chi.NewRouter()
// 	r := apiv1.Init()

// 	corsWrapper := cors.New(cors.Options{
// 		AllowedOrigins:   []string{"*"},
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
// 		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-CSRF-Token"},
// 		ExposedHeaders:   []string{"Content-Length", "Cache-Control", "Content-Language", "Content-Type", "Expires", "Last-Modified", "Pragma", "Authorization"},
// 		MaxAge:           86400,
// 		AllowCredentials: true,
// 		Debug:            false,
// 	})

// 	handler := corsWrapper.Handler(r)

// 	srv := &http.Server{
// 		Handler:           handler,
// 		Addr:              ":3001",
// 		ReadHeaderTimeout: 3 * time.Second,
// 		ReadTimeout:       5 * time.Second,
// 		WriteTimeout:      7 * time.Second,
// 		IdleTimeout:       30 * time.Second,
// 	}

// 	s := &Server{
// 		Server: srv,
// 		Router: r,
// 	}

// 	return s, nil
// }

// // Start runs the HTTP server
// func (s *Server) Start() {
// }

// // Stop stops the HTTP server
// func (s *Server) Stop() {
// }
