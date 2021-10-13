package urls

import (
	"context"
	"errors"
)

// Errors
var (
	ErrorUrlAlreadyExists = errors.New("URL is aslready exists")
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
		return EmptyUrlHash, ErrorUrlAlreadyExists
	}

	// append URL into repository
	if err = u.urlStore.Create(ctx, &Url{Hash: urlHash, URL: url}); err != nil {
		return EmptyUrlHash, err
	}

	return urlHash, nil
}
