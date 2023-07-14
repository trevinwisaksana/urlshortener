package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	editUsecase "github.com/trevinwisaksana/trevin-urlshortener/usecase/edit"
	loginUsecase "github.com/trevinwisaksana/trevin-urlshortener/usecase/login"
	registerUsecase "github.com/trevinwisaksana/trevin-urlshortener/usecase/register"
	shortenerUsecase "github.com/trevinwisaksana/trevin-urlshortener/usecase/shorten"
)

type EditUrlRequest struct {
	Username         string `json:"username" binding:"required"`
	CurrentShortlink string `json:"current_shortlink" binding:"required"`
	NewShortLink     string `json:"new_shortlink" binding:"required,alphanum,max=5"`
}

func MakeEditUrlEndpoint(u editUsecase.UrlShortenerEditUsecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*EditUrlRequest)
		return u.EditURL(ctx, req.Username, req.CurrentShortlink, req.NewShortLink)
	}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

func MakeLoginEndpoint(u loginUsecase.UrlShortenerLoginUsecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*LoginRequest)
		return u.Login(ctx, req.Username, req.Password)
	}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func MakeRegisterEndpoint(u registerUsecase.UrlShortenerRegisterUsecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*RegisterRequest)
		return u.Register(ctx, req.Username, req.Password, req.FullName, req.Email)
	}
}

type ShortenURLRequest struct {
	Username string `json:"username" binding:"required"`
	LongUrl  string `json:"long_url" binding:"required"`
}

func MakeShortenUrlEndpoint(u shortenerUsecase.UrlShortenerUsecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*ShortenURLRequest)
		return u.CreateShortURL(ctx, req.Username, req.LongUrl)
	}
}
