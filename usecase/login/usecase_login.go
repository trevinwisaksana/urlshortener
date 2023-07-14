package usecase

import (
	"context"

	"github.com/trevinwisaksana/trevin-urlshortener/cmd/initialize/config"
	"github.com/trevinwisaksana/trevin-urlshortener/internal/token"
	"github.com/trevinwisaksana/trevin-urlshortener/model"
	urlshortenerrepo "github.com/trevinwisaksana/trevin-urlshortener/repository/urlshortener/postgresql"
	"github.com/trevinwisaksana/trevin-urlshortener/tools"
)

type UrlShortenerLoginUsecase interface {
	Login(ctx context.Context, username string, password string) (LoginUserResponse, error)
}

type urlShortenerLoginUsecase struct {
	store      urlshortenerrepo.PostgreSQL
	tokenMaker token.PasetoMaker
	config     config.Config
}

func NewUrlShortenerLoginUsecase(store urlshortenerrepo.PostgreSQL, tokenMaker token.PasetoMaker, config config.Config) UrlShortenerLoginUsecase {
	return &urlShortenerLoginUsecase{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}
}

type LoginUserResponse struct {
	AccessToken string             `json:"access_token"`
	User        model.UserResponse `json:"user"`
}

func (u *urlShortenerLoginUsecase) Login(ctx context.Context, username string, password string) (LoginUserResponse, error) {
	var response LoginUserResponse

	user, err := u.store.GetUser(ctx, username)
	if err != nil {
		return response, err
	}

	if err := tools.ComparePasswordHash(password, user.HashedPassword); err != nil {
		return response, err
	}

	token, err := u.tokenMaker.CreateToken(user.Username, u.config.AccessTokenDuration)
	if err != nil {
		return response, err
	}

	userResponse := model.UserResponse{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	response = LoginUserResponse{
		AccessToken: token,
		User:        userResponse,
	}

	return response, nil
}
