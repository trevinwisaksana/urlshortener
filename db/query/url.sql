-- name: CreateShortURL :one
INSERT INTO urls (
  id, 
  long_url,
  short_url
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetLongURL :one
SELECT long_url FROM urls
WHERE id = $1 LIMIT 1;