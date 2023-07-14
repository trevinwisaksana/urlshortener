package usecase

import (
	"context"

	"github.com/trevinwisaksana/trevin-urlshortener/cmd/initialize/config"
	urlshortenerrepo "github.com/trevinwisaksana/trevin-urlshortener/repository/urlshortener/postgresql"
	"github.com/trevinwisaksana/trevin-urlshortener/tools"
)

type UrlShortenerUsecase interface {
	CreateShortURL(ctx context.Context, username string, longUrl string) (string, error)
}

type urlShortenerUsecase struct {
	store  urlshortenerrepo.PostgreSQL
	config config.Config
}

func NewUrlShortenerUsecase(store urlshortenerrepo.PostgreSQL, config config.Config) UrlShortenerUsecase {
	return &urlShortenerUsecase{
		store:  store,
		config: config,
	}
}

func (u *urlShortenerUsecase) CreateShortURL(ctx context.Context, username string, longUrl string) (string, error) {
	var response string

	randomID := tools.RandomAlphanumericString(5)

	result, err := u.store.CreateShortUrl(ctx, username, longUrl, randomID)
	if err != nil {
		return response, err
	}

	response = u.config.BaseUrl + result.ShortUrl

	return response, nil
}
