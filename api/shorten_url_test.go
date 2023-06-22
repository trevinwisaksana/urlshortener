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
	"github.com/trevinwisaksana/trevin-urlshortener/token"
	"github.com/trevinwisaksana/trevin-urlshortener/tools"
)

type eqCreateShortURLParamsMatcher struct {
	arg db.CreateShortURLParams
}

func (e eqCreateShortURLParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateShortURLParams)

	if !ok {
		return false
	}

	e.arg.ShortUrl = arg.ShortUrl
	e.arg.Owner = arg.Owner

	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateShortURLParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v", e.arg)
}

func EqCreateShortURLParams(arg db.CreateShortURLParams) gomock.Matcher {
	return eqCreateShortURLParamsMatcher{arg}
}

func TestShortenURL(t *testing.T) {
	user, _ := randomUser(t)
	url := dummyURL(t, user.Username)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.PasetoMaker)
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"long_url": url.LongUrl,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.PasetoMaker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				arg := db.CreateShortURLParams{
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
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.PasetoMaker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
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
		{
			name: "NoAuthorization",
			body: gin.H{
				"long_url": url.LongUrl,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.PasetoMaker) {

			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateShortURL(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
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

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func dummyURL(t *testing.T, owner string) (url db.Url) {
	randomID := tools.RandomAlphanumericString(5)

	url = db.Url{
		LongUrl:   "https://www.notion.so/stockbit/Backend-Engineering-Challenge-Link-Shortener-82bf71375701427c9cdd54a10a775ba6?pvs=4",
		ShortUrl:  randomID,
		CreatedAt: time.Now(),
		Owner:     owner,
	}

	return
}
