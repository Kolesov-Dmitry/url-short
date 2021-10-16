package srv

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Server is a HTTP server
type Server struct {
	router *mux.Router
	server *http.Server
}

// NewServer makes new Server instance
// Input:
//   port is a TCP port to listen
func NewServer(port int) *Server {
	return &Server{
		router: mux.NewRouter(),
		server: &http.Server{
			Addr:              fmt.Sprintf(":%d", port),
			ReadTimeout:       30 * time.Second,
			WriteTimeout:      30 * time.Second,
			ReadHeaderTimeout: 30 * time.Second,
		},
	}
}

// RegisterService registers new HTTP service
func (s *Server) RegisterService(service Service) {
	service.Register(s.router)
}

// Start starts the port listening
func (s *Server) Start() error {
	s.router.Handle("/", http.FileServer(http.Dir("static")))
	s.server.Handler = s.router

	return s.server.ListenAndServe()
}

// Close gracefully shuts down http server
func (s *Server) Close(ctx context.Context) {
	s.server.Shutdown(ctx)
}
