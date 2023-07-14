package usecase

import (
	"context"
	"strings"

	"github.com/trevinwisaksana/trevin-urlshortener/cmd/initialize/config"
	"github.com/trevinwisaksana/trevin-urlshortener/model"
	urlshortenerrepo "github.com/trevinwisaksana/trevin-urlshortener/repository/urlshortener/postgresql"
)

type UrlShortenerEditUsecase interface {
	EditURL(ctx context.Context, username string, currentShortlink string, newShortlink string) (model.Url, error)
}

type urlShortenerEditUsecase struct {
	store  urlshortenerrepo.PostgreSQL
	config config.Config
}

func NewUrlShortenerEditUsecase(store urlshortenerrepo.PostgreSQL, config config.Config) UrlShortenerEditUsecase {
	return &urlShortenerEditUsecase{
		store:  store,
		config: config,
	}
}

func (u *urlShortenerEditUsecase) EditURL(ctx context.Context, username string, currentShortlink string, newShortlink string) (model.Url, error) {
	var result model.Url

	prevShortUrl := strings.TrimPrefix(currentShortlink, u.config.BaseUrl)

	if err := u.isValid(ctx, prevShortUrl, username); err != nil {
		return result, err
	}

	result, err := u.store.UpdateShortURL(ctx, prevShortUrl, newShortlink)

	return result, err
}

func (u *urlShortenerEditUsecase) isValid(ctx context.Context, currentShortlink string, username string) error {
	response, err := u.store.GetURL(ctx, currentShortlink)
	if err != nil {
		return err
	}

	if response.Owner != username {
		return err
	}

	return nil
}
