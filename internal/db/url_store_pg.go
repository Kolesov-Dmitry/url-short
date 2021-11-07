package db

import (
	"context"
	"database/sql"
	"errors"
	"url-short/internal/repos/urls"
)

// UrlStorePg is an UrlStore implementation above PostgreSql
type UrlStorePg struct {
	conn *sql.DB
}

// NewUrlStorePg creates new UrlStorePg instance
func NewUrlStorePg() *UrlStorePg {
	return &UrlStorePg{
		conn: nil,
	}
}

// Connect establishes connection to the database using provided connection URL
// Inputs:
//   dbUrl - database connection URL
// Output:
//   Returns error if failed
func (s *UrlStorePg) Connect(dbUrl string) error {
	var err error
	s.conn, err = sql.Open("pgx", dbUrl)

	return err
}

// Create writes new URL to the database
// Inputs:
//   ctx - operation context
//   u - URL to write to the database
// Output:
//   Returns error if failed
func (s *UrlStorePg) Create(ctx context.Context, u *urls.Url) error {
	_, err := s.conn.ExecContext(ctx, "INSERT INTO urls.urls (url_hash, url) VALUES ($1, $2)", string(u.Hash), u.URL)
	return err
}

// Close shuts down database connection
func (s *UrlStorePg) Close() {
	s.conn.Close()
}

// Read fetches URL from the database by given URL hash
// Inputs:
//   ctx - operation context
//   hash - URL hash
// Output:
//   Returns found URL if succeeded, otherwise returns error
func (s *UrlStorePg) Read(ctx context.Context, hash urls.UrlHash) (*urls.Url, error) {
	var url string

	err := s.conn.QueryRowContext(ctx, "SELECT url FROM urls.urls WHERE url_hash=$1", string(hash)).Scan(&url)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &urls.Url{Hash: hash, URL: url}, nil
}
