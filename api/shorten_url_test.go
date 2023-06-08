package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/trevinwisaksana/trevin-urlshortener/db/mock"
	db "github.com/trevinwisaksana/trevin-urlshortener/db/sqlc"
)

type eqCreateShortURLParamsMatcher struct {
	arg db.CreateShortURLParams
}

func (e eqCreateShortURLParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateShortURLParams)

	if !ok {
		return false
	}

	e.arg.ID = arg.ID
	e.arg.ShortUrl = arg.ShortUrl

	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateShortURLParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v", e.arg)
}

func EqCreateShortURLParams(arg db.CreateShortURLParams) gomock.Matcher {
	return eqCreateShortURLParamsMatcher{arg}
}

func TestShortenURL(t *testing.T) {
	url := dummyURL(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"long_url": url.LongUrl,
			},
			buildStub: func(store *mockdb.MockStore) {
				arg := db.CreateShortURLParams{
					ID:       url.ID,
					LongUrl:  url.LongUrl,
					ShortUrl: url.ShortUrl,
				}

				store.EXPECT().
					CreateShortURL(gomock.Any(), EqCreateShortURLParams(arg)).
					Times(1).
					Return(url, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"long_url": url.LongUrl,
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateShortURL(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Url{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStub(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshall body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/shortener"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func dummyURL(t *testing.T) (url db.Url) {
	url = db.Url{
		ID:        "1",
		LongUrl:   "https://www.notion.so/stockbit/Backend-Engineering-Challenge-Link-Shortener-82bf71375701427c9cdd54a10a775ba6?pvs=4",
		ShortUrl:  "http://localhost:8080/1d0gd",
		CreatedAt: time.Now(),
	}

	return
}
