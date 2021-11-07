package db

import (
	"context"
	"log"
	"time"
	"url-short/internal/repos/urls"

	"github.com/go-redis/redis/v8"
)

// UrlCacheRedis is urls cache. Used Redis as a storage
// Implements UrlStore interface
type UrlCacheRedis struct {
	client *redis.Client
	store  urls.UrlStore
}

const (
	expirePeriod time.Duration = 10 * time.Minute
)

// NewUrlCacheRedis makes new UrlCacheRedis instance above provided persistent storage
// Inputs:
//    - store is a persistent URLs storage
//    - addr is a redis server address
//    - password is a password from the redis server
func NewUrlCacheRedis(store urls.UrlStore, addr string, password string) *UrlCacheRedis {
	return &UrlCacheRedis{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		}),
		store: store,
	}
}

// Create prompts create URL request to the percistent storage
func (c *UrlCacheRedis) Create(ctx context.Context, u *urls.Url) error {
	return c.store.Create(ctx, u)
}

// Read prompts create URL request to the percistent storage
func (c *UrlCacheRedis) Read(ctx context.Context, hash urls.UrlHash) (*urls.Url, error) {
	if val, err := c.client.Get(ctx, string(hash)).Result(); err == nil {

		go func(key string) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if _, err := c.client.Expire(ctx, key, expirePeriod).Result(); err != nil {
				log.Println(err)
			}
		}(string(hash))

		return &urls.Url{
			Hash: hash,
			URL:  val,
		}, nil
	}

	url, err := c.store.Read(ctx, hash)
	if err != nil {
		return nil, err
	}

	if url != nil {
		go func(key string, url string) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if _, err := c.client.Set(ctx, key, url, expirePeriod).Result(); err != nil {
				log.Println(err)
			}
		}(string(url.Hash), url.URL)
	}

	return url, nil
}

// Close shuts down redis connection
func (c *UrlCacheRedis) Close() {
	c.client.Close()
}
