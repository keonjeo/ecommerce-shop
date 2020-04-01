package app

import (
	"github.com/dankobgd/ecommerce-shop/config"
	"github.com/dankobgd/ecommerce-shop/zlog"
)

// SetConfig option for the app
func SetConfig(cfg *config.Config) Option {
	return func(a *App) error {
		a.cfg = cfg
		return nil
	}
}

// SetLogger option for the app
func SetLogger(logger *zlog.Logger) Option {
	return func(a *App) error {
		a.log = logger
		return nil
	}
}

// SetServer option for the app
func SetServer(server *Server) Option {
	return func(a *App) error {
		a.srv = server
		return nil
	}
}
