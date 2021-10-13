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
	Read(ctx context.Context, hash UrlHash) (*UrlLinking, error)
}

/*
type Users struct {
	ustore UserStore
}

func NewUsers(ustore UserStore) *Users {
	return &Users{
		ustore: ustore,
	}
}

func (us *Users) Create(ctx context.Context, u User) (*User, error) {
	u.ID = uuid.New()
	id, err := us.ustore.Create(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}
	u.ID = *id
	return &u, nil
}

func (us *Users) Read(ctx context.Context, uid uuid.UUID) (*User, error) {
	u, err := us.ustore.Read(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("read user error: %w", err)
	}
	return u, nil
}

func (us *Users) Delete(ctx context.Context, uid uuid.UUID) (*User, error) {
	u, err := us.ustore.Read(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("search user error: %w", err)
	}
	return u, us.ustore.Delete(ctx, uid)
}
*/
