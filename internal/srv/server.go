package srv

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Server is a HTTP server
type Server struct {
	port   int
	router *mux.Router
}

// NewServer makes new Server instance
// Input:
//   port is a TCP port to listen
func NewServer(port int) *Server {
	return &Server{
		port:   port,
		router: mux.NewRouter(),
	}
}

// Start starts the port listening
func (s *Server) Start() error {
	fs := http.FileServer(http.Dir("static"))
	s.router.Handle("/static/", http.StripPrefix("/static/", fs))

	return http.ListenAndServe(
		fmt.Sprintf(":%d", s.port),
		s.router,
	)
}
