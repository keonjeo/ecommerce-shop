package app

// App represents the app struct
type App struct {
	srv *Server
	// log *log.Logger
	// t goi18n.TranslateFunc
}

type appOption func(*App)

// New creates the new App
func New(options ...appOption) *App {
	app := &App{}

	for _, option := range options {
		option(app)
	}

	return app
}

// Srv ...
func (a *App) Srv() *Server {
	return a.srv
}

// func (a *App) Log() *log.Logger {
// 	return a.log
// }

// func (a *App) T(translationID string, args ...interface{}) string {
// 	return a.t(translationID, args...)
// }
