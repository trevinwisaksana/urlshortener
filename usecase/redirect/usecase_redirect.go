package usecase

import (
	"context"

	"github.com/trevinwisaksana/trevin-urlshortener/model"
	urlshortenerrepo "github.com/trevinwisaksana/trevin-urlshortener/repository/urlshortener/postgresql"
)

type UrlShortenerRedirectUsecase interface {
	GetURL(ctx context.Context, shortlinkID string) (model.Url, error)
}

type urlShortenerRedirectUsecase struct {
	store urlshortenerrepo.PostgreSQL
}

func NewUrlShortenerRedirectUsecase(store urlshortenerrepo.PostgreSQL) UrlShortenerRedirectUsecase {
	return &urlShortenerRedirectUsecase{
		store: store,
	}
}

func (u *urlShortenerRedirectUsecase) GetURL(ctx context.Context, shortlinkID string) (model.Url, error) {
	result, err := u.store.GetURL(ctx, shortlinkID)
	if err != nil {
		return result, err
	}

	return result, nil
}
