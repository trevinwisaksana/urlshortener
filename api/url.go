package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/trevinwisaksana/trevin-urlshortener/db/sqlc"
	random "github.com/trevinwisaksana/trevin-urlshortener/tools"
)

type ShortenURLRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func (server *Server) createShortURL(ctx *gin.Context) {
	var req ShortenURLRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	randomID := random.RandomString(5)

	arg := db.CreateShortURLParams{
		ID:       randomID,
		LongUrl:  req.LongUrl,
		ShortUrl: randomID,
	}

	result, err := server.store.CreateShortURL(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Redirect(http.StatusFound, result)
}
