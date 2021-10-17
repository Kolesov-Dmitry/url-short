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
	if errors.Is(err, urls.ErrUrlAlreadyExists) {
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
	if _, err := rw.Write([]byte(data)); err != nil {
		log.Println(err)
	}
}

// expandUrlGetHandler '/{urlHash:[0-9]+}' GET request handler
func (s *ShortenService) expandUrlGetHandler(rw http.ResponseWriter, r *http.Request) {
	urlHash, ok := mux.Vars(r)["urlHash"]
	if !ok {
		http.Error(rw, "Wrong URL", http.StatusBadRequest)
		return
	}

	url, err := s.urlsRepo.FetchUrlByHash(r.Context(), urls.UrlHash(urlHash))
	if errors.Is(err, urls.ErrUrlNotFound) {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Oops...", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	rw.Header().Set("Location", url)
	rw.WriteHeader(http.StatusTemporaryRedirect)
}

// Register registers service handlers
// Inputs:
//   router - HTTP mux router
func (s *ShortenService) Register(router *mux.Router) {
	router.HandleFunc("/api/v1/shorten", s.shortenPostHandler).Methods(http.MethodPost)
	router.HandleFunc("/{urlHash:[0-9]+}", s.expandUrlGetHandler).Methods(http.MethodGet)
}
