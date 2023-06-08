// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: url.sql

package db

import (
	"context"
)

const createShortURL = `-- name: CreateShortURL :one
INSERT INTO urls (
  id, 
  long_url,
  short_url
) VALUES (
  $1, $2, $3
) RETURNING id, long_url, short_url, created_at
`

type CreateShortURLParams struct {
	ID       string `json:"id"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

func (q *Queries) CreateShortURL(ctx context.Context, arg CreateShortURLParams) (Url, error) {
	row := q.db.QueryRowContext(ctx, createShortURL, arg.ID, arg.LongUrl, arg.ShortUrl)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.LongUrl,
		&i.ShortUrl,
		&i.CreatedAt,
	)
	return i, err
}

const getLongURL = `-- name: GetLongURL :one
SELECT long_url FROM urls
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetLongURL(ctx context.Context, id string) (string, error) {
	row := q.db.QueryRowContext(ctx, getLongURL, id)
	var long_url string
	err := row.Scan(&long_url)
	return long_url, err
}