package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
}

// NewServer ...
func NewServer() (*Server, error) {
	// r := chi.NewRouter()
	r := apiv1.Init()

	corsWrapper := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Content-Length", "Cache-Control", "Content-Language", "Content-Type", "Expires", "Last-Modified", "Pragma", "Authorization"},
		MaxAge:           86400,
		AllowCredentials: true,
		Debug:            false,
	})

	handler := corsWrapper.Handler(r)

	srv := &http.Server{
		Handler:           handler,
		Addr:              ":3001",
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      7 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	s := &Server{
		Server: srv,
		Router: r,
	}

	return s, nil
}

// Start runs the HTTP server
func (s *Server) Start() {
	fmt.Println("start")
	done := make(chan struct{})
	defer close(done)

	go func() {
		waitForTermination(done)
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		s.Stop(ctx)
	}()

	if err := s.Server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("Server listen failed")
		// log.WithError(err).Fatal("http server listen failed")
	}

}

// Stop stops the HTTP server
func (s *Server) Stop(ctx context.Context) {
	fmt.Println("stop")
	s.Server.Shutdown(ctx)
}

// WaitForShutdown blocks until the system signals termination or done has a value
func waitForTermination(done <-chan struct{}) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-signals:
		log.Printf("Triggering shutdown from singlas %s", sig)
		// log.Infof("Triggering shutdown from signal %s", sig)
	case <-done:
		log.Printf("Shutting down")
		// log.Infof("Shutting down...")
	}
}
