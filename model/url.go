package model

import "time"

type ShortenURLRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

type ShortenURLResponse struct {
	ShortUrl string `json:"short_url"`
}

type RedirectURLRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

type Url struct {
	ID        int64     `json:"id"`
	LongUrl   string    `json:"long_url"`
	ShortUrl  string    `json:"short_url"`
	CreatedAt time.Time `json:"created_at"`
	Owner     string    `json:"owner"`
}
