package url

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"url-short/internal/repos/urls"

	"github.com/gorilla/mux"
)

// ShortenService implements srv.Service interface
type ShortenService struct {
	urlsRepo *urls.Urls
}

// NewService makes new url.Service instance
// Inputs:
//   us - UrlStore implementation
//   ls - LinkingStore implementation
// Output:
//   Returns new ShortenService instance
func NewService(us urls.UrlStore, ls urls.LinkingStore) *ShortenService {
	return &ShortenService{
		urlsRepo: urls.NewUrls(us, ls),
	}
}

// shortenPostHandler '/api/v1/shorten' POST request handler
func (s *ShortenService) shortenPostHandler(rw http.ResponseWriter, r *http.Request) {
	holder := struct {
		Url string `json:"url"`
	}{}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&holder); err != nil {
		http.Error(rw, "Failed to read request body", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	urlHash, err := s.urlsRepo.AppendUrl(r.Context(), holder.Url)
	if errors.Is(err, urls.ErrorUrlAlreadyExists) {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}

	if err != nil {
		http.Error(rw, "Oops...", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	answer := struct {
		Hash string `json:"hash"`
	}{}
	answer.Hash = string(urlHash)

	data, err := json.Marshal(&answer)
	if err != nil {
		http.Error(rw, "Oops...", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(data))
}

// Register registers service handlers
// Inputs:
//   router - HTTP mux router
func (s *ShortenService) Register(router *mux.Router) {
	router.HandleFunc("/api/v1/shorten", s.shortenPostHandler).Methods(http.MethodPost)
}
