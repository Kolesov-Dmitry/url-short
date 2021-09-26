package url

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Service implements srv.Service interface
type Service struct {
}

// NewService makes new url.Service instance
func NewService() *Service {
	return &Service{}
}

func (s *Service) shortenPostHandler(rw http.ResponseWriter, r *http.Request) {

}

func (s *Service) Register(router *mux.Router) {
	router.HandleFunc("api/v1/shorten", s.shortenPostHandler).Methods(http.MethodPost)
}
