package api

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/trevinwisaksana/trevin-urlshortener/db/sqlc"
	"github.com/trevinwisaksana/trevin-urlshortener/token"
	tools "github.com/trevinwisaksana/trevin-urlshortener/tools"
)

type ShortenURLRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

type ShortenURLResponse struct {
	ShortUrl string `json:"short_url"`
}

func newShortenURLResponse(baseUrl string, response db.Url) ShortenURLResponse {
	return ShortenURLResponse{
		ShortUrl: baseUrl + response.ShortUrl,
	}
}

func (server *Server) createShortURL(ctx *gin.Context) {
	var req ShortenURLRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorzationPayloadKey).(*token.Payload)

	randomID := tools.RandomString(5)

	arg := db.CreateShortURLParams{
		LongUrl:  req.LongUrl,
		ShortUrl: randomID,
		Owner:    authPayload.Username,
	}

	result, err := server.store.CreateShortURL(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := newShortenURLResponse(server.config.BaseUrl, result)

	ctx.JSON(http.StatusOK, response)
}

type RedirectURLRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

func (server *Server) redirectURL(ctx *gin.Context) {
	var req RedirectURLRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := server.store.GetLongURL(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.Redirect(http.StatusFound, result)
}

type EditURLRequest struct {
	CurrentShortlink string `json:"current_shortlink" binding:"required"`
	NewShortLink     string `json:"new_shortlink" binding:"required"`
}

func (server *Server) editURL(ctx *gin.Context) {
	var req EditURLRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorzationPayloadKey).(*token.Payload)

	prevShortUrl := strings.TrimPrefix(req.CurrentShortlink, server.config.BaseUrl)
	newShortlink := req.NewShortLink

	if errorCode, err := server.isValid(ctx, prevShortUrl, newShortlink, authPayload.Username); err != nil {
		ctx.JSON(errorCode, errorResponse(err))
		return
	}

	arg := db.UpdateShortURLParams{
		ShortUrl:        prevShortUrl,
		CustomShortlink: newShortlink,
	}

	result, err := server.store.UpdateShortURL(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusConflict, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) isValid(ctx *gin.Context, shortUrl string, newShortUrl string, username string) (int, error) {
	response, err := server.store.GetURL(ctx, shortUrl)
	if err != nil {
		return http.StatusBadRequest, err
	}

	if response.Owner != username {
		err := errors.New("User is not authorized to edit this url")
		return http.StatusUnauthorized, err
	}

	if len(newShortUrl) > 5 {
		err := errors.New("Shortlink cannot be longer than 5 characters")
		return http.StatusBadRequest, err
	}

	isAlphanumeric := regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(newShortUrl)

	if !isAlphanumeric {
		err := errors.New("Shortlink must only contain alphanumeric characters")
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}
