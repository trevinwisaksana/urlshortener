package shorten

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trevinwisaksana/trevin-urlshortener/internal/token"
	"github.com/trevinwisaksana/trevin-urlshortener/model/constant"
	usecase "github.com/trevinwisaksana/trevin-urlshortener/usecase/shorten"
)

type restServer struct {
	u usecase.UrlShortenerUsecase
}

type ShortenURLRequest struct {
	LongUrl string `json:"long_url"`
}

type ShortenURLResponse struct {
	ShortUrl string `json:"short_url"`
}

func (r *restServer) ShortenURL(ctx *gin.Context) {
	var req ShortenURLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(constant.AuthorzationPayloadKey).(*token.Payload)

	response, err := r.u.CreateShortURL(ctx, authPayload.Username, req.LongUrl)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result := ShortenURLResponse{
		ShortUrl: response,
	}

	ctx.JSON(http.StatusOK, result)
}
