package db

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
	"url-short/internal/repos/urls"

	"github.com/google/uuid"
)

const (
	dateTimeFormat = "2006-01-02 15:04:05"
)

// LinkingStorePg is an UrlLinking implementation above PostgreSql
type LinkingStorePg struct {
	conn *sql.DB
}

// NewLinkingStorePg creates new LinkingStorePg instance
func NewLinkingStorePg() *LinkingStorePg {
	return &LinkingStorePg{
		conn: nil,
	}
}

// Connect establishes connection to the database using provided connection URL
// Inputs:
//   dbUrl - database connection URL
// Output:
//   Returns error if failed
func (s *LinkingStorePg) Connect(dbUrl string) error {
	var err error
	s.conn, err = sql.Open("pgx", dbUrl)

	return err
}

// Create writes new URL linking to the database
// Inputs:
//   ctx - operation context
//   u - URL to write to the database
// Output:
//   Returns error if failed
func (s *LinkingStorePg) Create(ctx context.Context, urlHash urls.UrlHash) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	_, err = s.conn.ExecContext(ctx, "INSERT INTO urls.urls_linking (id, url_hash) VALUES ($1, $2)", id, string(urlHash))
	return err
}

// Close shuts down database connection
func (s *LinkingStorePg) Close() {
	s.conn.Close()
}

// Read fetches URL linking from the database by given URL hash
// Inputs:
//   ctx - operation context
//   hash - URL hash
// Output:
//   Returns found URL linkings if succeeded, otherwise returns error
func (s *LinkingStorePg) Read(ctx context.Context, hash urls.UrlHash) chan string {
	outChan := make(chan string, 100)

	go func() {
		defer close(outChan)

		rows, err := s.conn.QueryContext(ctx, "SELECT visited_at FROM urls.urls_linking WHERE url_hash=$1", string(hash))
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(err)
			return
		}
		defer rows.Close()

		var visited_at time.Time
		for rows.Next() {
			if err := rows.Scan(&visited_at); err != nil {
				log.Println(err)
				return
			}

			outChan <- visited_at.Format(dateTimeFormat)
		}
	}()

	return outChan
}
