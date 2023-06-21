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
  long_url,
  short_url,
  owner
) VALUES (
  $1, $2, $3
) RETURNING id, long_url, short_url, created_at, owner
`

type CreateShortURLParams struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
	Owner    string `json:"owner"`
}

func (q *Queries) CreateShortURL(ctx context.Context, arg CreateShortURLParams) (Url, error) {
	row := q.db.QueryRowContext(ctx, createShortURL, arg.LongUrl, arg.ShortUrl, arg.Owner)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.LongUrl,
		&i.ShortUrl,
		&i.CreatedAt,
		&i.Owner,
	)
	return i, err
}

const getLongURL = `-- name: GetLongURL :one
SELECT long_url FROM urls
WHERE short_url = $1 LIMIT 1
`

func (q *Queries) GetLongURL(ctx context.Context, shortUrl string) (string, error) {
	row := q.db.QueryRowContext(ctx, getLongURL, shortUrl)
	var long_url string
	err := row.Scan(&long_url)
	return long_url, err
}

const getURL = `-- name: GetURL :one
SELECT id, long_url, short_url, created_at, owner FROM urls
WHERE short_url = $1 LIMIT 1
`

func (q *Queries) GetURL(ctx context.Context, shortUrl string) (Url, error) {
	row := q.db.QueryRowContext(ctx, getURL, shortUrl)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.LongUrl,
		&i.ShortUrl,
		&i.CreatedAt,
		&i.Owner,
	)
	return i, err
}

const updateShortURL = `-- name: UpdateShortURL :one
UPDATE urls
SET short_url = $2
WHERE short_url = $1
RETURNING id, long_url, short_url, created_at, owner
`

type UpdateShortURLParams struct {
	ShortUrl        string `json:"short_url"`
	CustomShortlink string `json:"custom_shortlink"`
}

func (q *Queries) UpdateShortURL(ctx context.Context, arg UpdateShortURLParams) (Url, error) {
	row := q.db.QueryRowContext(ctx, updateShortURL, arg.ShortUrl, arg.CustomShortlink)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.LongUrl,
		&i.ShortUrl,
		&i.CreatedAt,
		&i.Owner,
	)
	return i, err
}
