package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
	mockdb "github.com/trevinwisaksana/trevin-urlshortener/db/mock"
	db "github.com/trevinwisaksana/trevin-urlshortener/db/sqlc"
	"github.com/trevinwisaksana/trevin-urlshortener/token"
	"github.com/trevinwisaksana/trevin-urlshortener/tools"
)

func TestEditURL(t *testing.T) {
	user, _ := randomUser(t)
	second_user, _ := randomUser(t)
	url := dummyURL(t, user.Username)
	updatedUrl := updatedDummyURL(t, url.ShortUrl, user.Username)

	t.Log(updatedUrl.ShortUrl)

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
				"current_shortlink": url.ShortUrl,
				"new_shortlink":     updatedUrl.ShortUrl,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.PasetoMaker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetURL(gomock.Any(), gomock.Eq(url.ShortUrl)).
					Times(1).
					Return(url, nil)

				updateUrlArg := db.UpdateShortURLParams{
					ShortUrl:        url.ShortUrl,
					CustomShortlink: updatedUrl.ShortUrl,
				}

				store.EXPECT().
					UpdateShortURL(gomock.Any(), gomock.Eq(updateUrlArg)).
					Times(1).
					Return(updatedUrl, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "NoAuthorization",
			body: gin.H{
				"current_shortlink": url.ShortUrl,
				"new_shortlink":     updatedUrl.ShortUrl,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.PasetoMaker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, second_user.Username, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetURL(gomock.Any(), gomock.Eq(url.ShortUrl)).
					Times(1).
					Return(url, nil)

				updateUrlArg := db.UpdateShortURLParams{
					ShortUrl:        url.ShortUrl,
					CustomShortlink: updatedUrl.ShortUrl,
				}

				store.EXPECT().
					UpdateShortURL(gomock.Any(), gomock.Eq(updateUrlArg)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "NonUniqueShortlink",
			body: gin.H{
				"current_shortlink": url.ShortUrl,
				"new_shortlink":     updatedUrl.ShortUrl,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.PasetoMaker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetURL(gomock.Any(), gomock.Eq(url.ShortUrl)).
					Times(1).
					Return(url, nil)

				updateUrlArg := db.UpdateShortURLParams{
					ShortUrl:        url.ShortUrl,
					CustomShortlink: updatedUrl.ShortUrl,
				}

				store.EXPECT().
					UpdateShortURL(gomock.Any(), gomock.Eq(updateUrlArg)).
					Times(1).
					Return(db.Url{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusConflict, recorder.Code)
			},
		},
		{
			name: "ShortlinkIsTooLong",
			body: gin.H{
				"current_shortlink": url.ShortUrl,
				"new_shortlink":     "long-shortlink",
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.PasetoMaker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetURL(gomock.Any(), gomock.Eq(url.ShortUrl)).
					Times(0)

				updateUrlArg := db.UpdateShortURLParams{
					ShortUrl:        url.ShortUrl,
					CustomShortlink: "long-shortlink",
				}

				store.EXPECT().
					UpdateShortURL(gomock.Any(), gomock.Eq(updateUrlArg)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "NonAlphanumeric",
			body: gin.H{
				"current_shortlink": url.ShortUrl,
				"new_shortlink":     "@@@@",
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.PasetoMaker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetURL(gomock.Any(), gomock.Eq(url.ShortUrl)).
					Times(0)

				updateUrlArg := db.UpdateShortURLParams{
					ShortUrl:        url.ShortUrl,
					CustomShortlink: "@@@@",
				}

				store.EXPECT().
					UpdateShortURL(gomock.Any(), gomock.Eq(updateUrlArg)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
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

			url := "/edit"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func updatedDummyURL(t *testing.T, id string, owner string) (url db.Url) {
	randomID := tools.RandomAlphanumericString(5)

	url = db.Url{
		LongUrl:   "https://www.notion.so/stockbit/Backend-Engineering-Challenge-Link-Shortener-82bf71375701427c9cdd54a10a775ba6?pvs=4",
		ShortUrl:  randomID,
		CreatedAt: time.Now(),
		Owner:     owner,
	}

	return
}
