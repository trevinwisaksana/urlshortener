package usecase

import (
	"context"

	"github.com/trevinwisaksana/trevin-urlshortener/model"
	urlshortenerrepo "github.com/trevinwisaksana/trevin-urlshortener/repository/urlshortenerrepo/postgresql"
	"github.com/trevinwisaksana/trevin-urlshortener/tools"
)

type UrlShortenerRegisterUsecase interface {
	Register(ctx context.Context, username string, password string, fullName string, email string) (model.UserResponse, error)
}

type urlShortenerRegisterUsecase struct {
	store urlshortenerrepo.PostgreSQL
}

func NewUrlShortenerRegisterUsecase(store urlshortenerrepo.PostgreSQL) UrlShortenerRegisterUsecase {
	return &urlShortenerRegisterUsecase{
		store: store,
	}
}

func (u *urlShortenerRegisterUsecase) Register(ctx context.Context, username string, password string, fullName string, email string) (model.UserResponse, error) {
	var response model.UserResponse

	hashedPassword, err := tools.HashPassword(password)
	if err != nil {
		return response, err
	}

	user, err := u.store.CreateUser(ctx, username, hashedPassword, fullName, email)
	if err != nil {
		return response, err
	}

	response = model.UserResponse{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return response, nil
}
