package app

import (
	"github.com/dankobgd/ecommerce-shop/config"
	"github.com/dankobgd/ecommerce-shop/zlog"
)

// App represents the app struct
type App struct {
	srv *Server
	cfg *config.Config
	log *zlog.Logger
	// t goi18n.TranslateFunc
}

// Option for the app
type Option func(*App) error

// OptionCreator for the app
type OptionCreator func() []Option

// New creates the new App
func New(options ...Option) *App {
	app := &App{}

	for _, option := range options {
		option(app)
	}

	return app
}

// IsDev returs true if app is in development mode
func (a *App) IsDev() bool {
	return a.Cfg().ENV == "development"
}

// IsProd returs true if app is in production mode
func (a *App) IsProd() bool {
	return a.Cfg().ENV == "production"
}

// IsTest returs true if app is in test mode
func (a *App) IsTest() bool {
	return a.Cfg().ENV == "test"
}

// Srv ...
func (a *App) Srv() *Server {
	return a.srv
}

// SetServer ...
func (a *App) SetServer(srv *Server) {
	a.srv = srv
}

// Cfg ...
func (a *App) Cfg() *config.Config {
	return a.cfg
}

// SetConfig ...
func (a *App) SetConfig(cfg *config.Config) {
	a.cfg = cfg
}

// Log ...
func (a *App) Log() *zlog.Logger {
	return a.log
}

// SetLogger ...
func (a *App) SetLogger(logger *zlog.Logger) {
	a.log = logger
}

// func (a *App) T(translationID string, args ...interface{}) string {
// 	return a.t(translationID, args...)
// }
