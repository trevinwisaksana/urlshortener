package transport

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/trevinwisaksana/trevin-urlshortener/endpoint"
	usecase "github.com/trevinwisaksana/trevin-urlshortener/usecase/edit"
)

func MakeEditUrlHandler(u usecase.UrlShortenerEditUsecase) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/edit", httptransport.NewServer(
		endpoint.MakeEditUrlEndpoint(u),
		decodeEditUrlRequest,
		encodeResponse,
	))

	return mux
}

func decodeEditUrlRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.EditUrlRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
