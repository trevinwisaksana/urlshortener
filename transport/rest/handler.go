package transport

import (
	"context"
	"net/http"

	"github.com/trevinwisaksana/trevin-urlshortener/internal/token"
	"github.com/trevinwisaksana/trevin-urlshortener/model/constant"
)

type restShortenerServer struct {
	edit    http.Handler
	shorten http.Handler
}

type restRedirectServer struct {
	redirect http.Handler
}

type restRegistrationServer struct {
	register http.Handler
}

type restAuthServer struct {
	login http.Handler
}

func (r *restShortenerServer) EditURL(ctx context.Context, request *http.Request) {
	var req EditURLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	r.edit.ServeHTTP()

	authPayload := ctx.MustGet(constant.AuthorzationPayloadKey).(*token.Payload)

	response, err := r.editUsecase.EditURL(ctx.Request.Context(), authPayload.Username, req.CurrentShortlink, req.NewShortLink)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (r *restAuthServer) Login(ctx context.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := r.u.Login(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
