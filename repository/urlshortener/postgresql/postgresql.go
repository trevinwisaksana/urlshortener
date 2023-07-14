package urlshortenerrepo

import (
	"context"
	"database/sql"

	"github.com/trevinwisaksana/trevin-urlshortener/model"
	"github.com/trevinwisaksana/trevin-urlshortener/repository/urlshortener/postgresql/internal/sqlc"
)

//go:generate mockgen -source=repository/urlshortenerrepo/postgresql/postgresql.go -destination=repository/urlshortenerrepo/postgresql/mock/postgresql.go -package=mockurlshortener -self_package=github.com/trevinwisaksana/trevin-urlshortener
type PostgreSQL interface {
	UpdateShortURL(ctx context.Context, currentShortlink string, newShortlink string) (model.Url, error)
	GetURL(ctx context.Context, currentShortlink string) (model.Url, error)
	GetUser(ctx context.Context, username string) (model.User, error)
	CreateShortUrl(ctx context.Context, username string, longUrl string, randomID string) (model.Url, error)
	CreateUser(ctx context.Context, username string, hashedPassword string, fullName string, email string) (model.User, error)
}

// Store provides all functions to execute db queries and transactions
type postgresql struct {
	*sqlc.Queries
}

// NewPostgreSQL creates a new store
func NewPostgreSQL(db *sql.DB) PostgreSQL {
	return &postgresql{
		Queries: sqlc.New(db),
	}
}

func (p *postgresql) UpdateShortURL(ctx context.Context, currentShortlink string, newShortlink string) (model.Url, error) {
	arg := sqlc.UpdateShortURLParams{
		ShortUrl:        currentShortlink,
		CustomShortlink: newShortlink,
	}

	response, err := p.Queries.UpdateShortURL(ctx, arg)

	result := model.Url{
		ID:        response.ID,
		LongUrl:   response.LongUrl,
		ShortUrl:  response.ShortUrl,
		CreatedAt: response.CreatedAt,
		Owner:     response.Owner,
	}

	return result, err
}

func (p *postgresql) GetURL(ctx context.Context, currentShortlink string) (model.Url, error) {
	response, err := p.Queries.GetURL(ctx, currentShortlink)

	result := model.Url{
		ID:        response.ID,
		LongUrl:   response.LongUrl,
		ShortUrl:  response.ShortUrl,
		CreatedAt: response.CreatedAt,
		Owner:     response.Owner,
	}

	return result, err
}

func (p *postgresql) GetUser(ctx context.Context, username string) (model.User, error) {
	response, err := p.Queries.GetUser(ctx, username)

	result := model.User{
		Username:       response.Username,
		HashedPassword: response.HashedPassword,
		FullName:       response.FullName,
		Email:          response.Email,
		CreatedAt:      response.CreatedAt,
	}

	return result, err
}

func (p *postgresql) CreateShortUrl(ctx context.Context, username string, longUrl string, randomID string) (model.Url, error) {
	arg := sqlc.CreateShortURLParams{
		LongUrl:  longUrl,
		ShortUrl: randomID,
		Owner:    username,
	}

	response, err := p.Queries.CreateShortURL(ctx, arg)

	result := model.Url{
		ID:        response.ID,
		LongUrl:   response.LongUrl,
		ShortUrl:  response.ShortUrl,
		CreatedAt: response.CreatedAt,
		Owner:     response.Owner,
	}

	return result, err
}

func (p *postgresql) CreateUser(ctx context.Context, username string, hashedPassword string, fullName string, email string) (model.User, error) {
	arg := sqlc.CreateUserParams{
		Username:       username,
		HashedPassword: hashedPassword,
		FullName:       fullName,
		Email:          email,
	}

	response, err := p.Queries.CreateUser(ctx, arg)

	result := model.User{
		Username:       response.Username,
		HashedPassword: response.HashedPassword,
		FullName:       response.FullName,
		Email:          response.Email,
	}

	return result, err
}
