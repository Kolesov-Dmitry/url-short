package app

import (
	"context"
	"log"
	"url-short/internal/db"
	"url-short/internal/srv"
	"url-short/internal/url"
)

// Application manages url-short application control flow and main settings
type Application struct {
	settings       *settings
	server         *srv.Server
	shortenService *url.ShortenService
	urlStore       *db.UrlStorePg
	urlCache       *db.UrlCacheRedis
	linkingStore   *db.LinkingStorePg
}

// NewApplication creates new Application instance
// Fails if unable to read settings
func NewApplication() Application {
	return Application{
		settings:       nil,
		server:         nil,
		shortenService: nil,
		urlStore:       nil,
		urlCache:       nil,
		linkingStore:   nil,
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

	// Creating UrlStore
	a.urlStore = db.NewUrlStorePg()
	if err := a.urlStore.Connect(a.settings.dbURL); err != nil {
		return err
	}

	// Creating LinkingStore
	a.linkingStore = db.NewLinkingStorePg()
	if err := a.linkingStore.Connect(a.settings.dbURL); err != nil {
		return err
	}

	// Creating UrlCache
	a.urlCache = db.NewUrlCacheRedis(a.urlStore, a.settings.redisAddr, a.settings.redisPassword)

	// Creating HTTP server
	a.server = srv.NewServer(a.settings.port)
	a.registerServices()

	log.Println("Server is started")

	go func(server *srv.Server) {
		if err := server.Start(); err != nil {
			log.Fatal(err)
		}
	}(a.server)

	return nil
}

// Close gracefully shuts down HTTP server and database connections
func (a *Application) Close(ctx context.Context) {
	a.server.Close(ctx)
	a.urlCache.Close()
	a.urlStore.Close()
	a.linkingStore.Close()
}

// registerServices
func (a *Application) registerServices() {
	a.shortenService = url.NewService(a.urlCache, a.linkingStore)
	a.server.RegisterService(a.shortenService)
}
