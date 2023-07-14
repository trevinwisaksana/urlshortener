package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/trevinwisaksana/trevin-urlshortener/config"
	"github.com/trevinwisaksana/trevin-urlshortener/model"
	mockurlshortener "github.com/trevinwisaksana/trevin-urlshortener/repository/urlshortener/postgresql/mock"
	"github.com/trevinwisaksana/trevin-urlshortener/tools"
	urltools "github.com/trevinwisaksana/trevin-urlshortener/tools/url"
	usertools "github.com/trevinwisaksana/trevin-urlshortener/tools/user"
)

func TestEditURL(t *testing.T) {
	username := usertools.RandomUsername()
	url := urltools.DummyURL(username)
	updatedUrl := updatedDummyURL(t, url.ShortUrl, username)

	testCases := []struct {
		name             string
		username         string
		currentShortlink string
		newShortlink     string
		buildStub        func(store *mockurlshortener.MockPostgreSQL)
		checkResult      func(result model.Url, err error)
	}{
		{
			name:             "OK",
			username:         username,
			currentShortlink: url.ShortUrl,
			newShortlink:     updatedUrl.ShortUrl,
			buildStub: func(store *mockurlshortener.MockPostgreSQL) {
				store.EXPECT().
					GetURL(gomock.Any(), gomock.Eq(url.ShortUrl)).
					Times(1).
					Return(url, nil)

				store.EXPECT().
					UpdateShortURL(gomock.Any(), gomock.Eq(url.ShortUrl), gomock.Eq(updatedUrl.ShortUrl)).
					Times(1).
					Return(updatedUrl, nil)
			},
			checkResult: func(result model.Url, err error) {
				require.NoError(t, err)
				require.NotNil(t, result)

				require.Equal(t, result.LongUrl, url.LongUrl)
				require.Equal(t, result.ShortUrl, updatedUrl.ShortUrl)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockurlshortener.NewMockPostgreSQL(ctrl)
			tc.buildStub(store)

			config := config.Config{}

			usecase := NewUrlShortenerEditUsecase(store, config)
			response, err := usecase.EditURL(context.TODO(), tc.username, tc.currentShortlink, tc.newShortlink)
			tc.checkResult(response, err)
		})
	}
}

func updatedDummyURL(t *testing.T, id string, owner string) (url model.Url) {
	randomID := tools.RandomAlphanumericString(5)

	url = model.Url{
		LongUrl:   "https://www.notion.so/stockbit/Backend-Engineering-Challenge-Link-Shortener-82bf71375701427c9cdd54a10a775ba6?pvs=4",
		ShortUrl:  randomID,
		CreatedAt: time.Now(),
		Owner:     owner,
	}

	return
}
