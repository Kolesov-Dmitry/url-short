package urls

import (
	"context"
	"time"
)

// Url is an URL model type
type Url struct {
	Hash UrlHash
	URL  string
}

// UrlLinking holds an URL linking list
type UrlLinking struct {
	URL      string
	Linkings []time.Time
}

// UrlStore is an abstraction above urls storage
type UrlStore interface {
	Create(ctx context.Context, u *Url) error
	Read(ctx context.Context, hash UrlHash) (*Url, error)
}

// LinkingStore is an abstraction above url linkings storage
type LinkingStore interface {
	Create(ctx context.Context, hash UrlHash) error
	Read(ctx context.Context, hash UrlHash) chan string
}
