// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"time"
)

type Url struct {
	ID        string    `json:"id"`
	LongUrl   string    `json:"long_url"`
	ShortUrl  string    `json:"short_url"`
	CreatedAt time.Time `json:"created_at"`
}
