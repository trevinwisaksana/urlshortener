package transport

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/trevinwisaksana/trevin-urlshortener/endpoint"
	usecase "github.com/trevinwisaksana/trevin-urlshortener/usecase/login"
)

func MakeLoginHandler(u usecase.UrlShortenerLoginUsecase) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/login", httptransport.NewServer(
		endpoint.MakeLoginEndpoint(u),
		decodeLoginRequest,
		encodeResponse,
	))

	return mux
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.LoginUserRequest
	return req, nil
}

func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
