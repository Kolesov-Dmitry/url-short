package urls

import (
	"context"
	"errors"
)

// Errors
var (
	ErrUrlAlreadyExists = errors.New("URL is already exists")
	ErrUrlNotFound      = errors.New("URL was not found")
)

// Urls is an urls repository
type Urls struct {
	urlStore     UrlStore
	linkingStore LinkingStore
}

// NewUrls creates new Urls repository instance
// Inputs:
//   us - UrlStore implementation
//   ls - LinkingStore implementation
func NewUrls(us UrlStore, ls LinkingStore) *Urls {
	return &Urls{
		urlStore:     us,
		linkingStore: ls,
	}
}

// AppendUrl appends new URL into the repository
// Inputs:
//   ctx - context
//   url - URL to append into the repository
// Output:
//   Returns URL hash in case of success, otherwise returns error
func (u *Urls) AppendUrl(ctx context.Context, url string) (UrlHash, error) {
	urlHash := Hash(url)

	// check if provided URL already exists in the repository
	surl, err := u.urlStore.Read(ctx, urlHash)
	if err != nil {
		return EmptyUrlHash, err
	}

	if surl != nil {
		return EmptyUrlHash, ErrUrlAlreadyExists
	}

	// append URL into repository
	if err = u.urlStore.Create(ctx, &Url{Hash: urlHash, URL: url}); err != nil {
		return EmptyUrlHash, err
	}

	return urlHash, nil
}

// FetchUrlByHash fetches URL from the repository by given URL hash
// Inputs:
//   ctx  - context
//   hash - URL hash
// Output:
//   Returns found URL if succeeded, otherwise returns error
func (u *Urls) FetchUrlByHash(ctx context.Context, hash UrlHash) (string, error) {
	url, err := u.urlStore.Read(ctx, hash)
	if err != nil {
		return "", err
	}

	if url == nil {
		return "", ErrUrlNotFound
	}

	return url.URL, nil
}

// SaveUrlLinking saves URL linking statistics
// Inputs:
//   ctx  - context
//   hash - URL hash
// Output:
//   Returns error if failed
func (u *Urls) SaveUrlLinking(ctx context.Context, hash UrlHash) error {
	return u.linkingStore.Create(ctx, hash)
}
