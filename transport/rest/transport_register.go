package transport

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	usecase "github.com/trevinwisaksana/trevin-urlshortener/usecase/register"
)

func MakeRegisterHandler(u usecase.UrlShortenerRegisterUsecase) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/register", httptransport.NewServer(
		endpoint.MakeRegisterEndpoint(u),
		decodeLoginRequest,
		encodeResponse,
	))

	return mux
}

func decodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.LoginUserRequest
	return req, nil
}

func encodeRegisterResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
