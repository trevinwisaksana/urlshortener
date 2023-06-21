-- name: CreateShortURL :one
INSERT INTO urls (
  long_url,
  short_url,
  owner
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetURL :one
SELECT * FROM urls
WHERE short_url = $1 LIMIT 1;

-- name: GetLongURL :one
SELECT long_url FROM urls
WHERE short_url = $1 LIMIT 1;

-- name: UpdateShortURL :one
UPDATE urls
SET short_url = sqlc.arg(custom_shortlink)
WHERE short_url = $1
RETURNING *;