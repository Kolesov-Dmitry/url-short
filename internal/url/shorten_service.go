package url

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// ShortenService implements srv.Service interface
type ShortenService struct {
}

// NewService makes new url.Service instance
func NewService() *ShortenService {
	return &ShortenService{}
}

// shortenPostHandler '/api/v1/shorten' POST request handler
func (s *ShortenService) shortenPostHandler(rw http.ResponseWriter, r *http.Request) {
	holder := struct {
		Url string `json:"url"`
	}{}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&holder); err != nil {
		http.Error(rw, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/pain")
	rw.Write([]byte(holder.Url))
}

func (s *ShortenService) Register(router *mux.Router) {
	router.HandleFunc("/api/v1/shorten", s.shortenPostHandler).Methods(http.MethodPost)
}
