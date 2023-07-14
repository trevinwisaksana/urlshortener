package redirect

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/trevinwisaksana/trevin-urlshortener/usecase/redirect"
)

type restServer struct {
	u usecase.UrlShortenerRedirectUsecase
}

type RedirectURLRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

func (r *restServer) RedirectURL(ctx *gin.Context) {
	var req RedirectURLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := r.u.GetURL(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.Redirect(http.StatusFound, response.LongUrl)
}
