package app

import "url-short/internal/srv"

// Application manages url-short application control flow and main settings
type Application struct {
	settings *settings
	server   *srv.Server
}

// NewApplication creates new Application instance
// Fails if unable to read settings
func NewApplication() Application {
	return Application{
		settings: nil,
		server:   nil,
	}
}

// Run starts application main loop
// Returns error if failed
func (a *Application) Run() error {
	var err error
	a.settings, err = loadSettings()
	if err != nil {
		return err
	}

	a.server = srv.NewServer(a.settings.port)

	return a.server.Start()
}
