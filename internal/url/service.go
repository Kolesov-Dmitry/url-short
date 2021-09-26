package url

import (
	"encoding/json"
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

// shortenPostHandler '/api/v1/shorten' POST request handler
func (s *Service) shortenPostHandler(rw http.ResponseWriter, r *http.Request) {
	holder := struct {
		Url string `json:"url"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&holder); err != nil {
		http.Error(rw, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/pain")
	rw.Write([]byte(holder.Url))
}

func (s *Service) Register(router *mux.Router) {
	router.HandleFunc("/api/v1/shorten", s.shortenPostHandler).Methods(http.MethodPost)
}
