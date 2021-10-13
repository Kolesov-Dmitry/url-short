package app

import (
	"url-short/internal/srv"
	"url-short/internal/url"
)

// Application manages url-short application control flow and main settings
type Application struct {
	settings       *settings
	server         *srv.Server
	shortenService *url.ShortenService
}

// NewApplication creates new Application instance
// Fails if unable to read settings
func NewApplication() Application {
	return Application{
		settings:       nil,
		server:         nil,
		shortenService: nil,
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
	a.registerServices()

	return a.server.Start()
}

// registerServices
func (a *Application) registerServices() {
	a.shortenService = url.NewService()
	a.server.RegisterService(a.shortenService)
}
