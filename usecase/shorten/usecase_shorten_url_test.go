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
	usertools "github.com/trevinwisaksana/trevin-urlshortener/tools/user"
)

func TestShortenURL(t *testing.T) {
	user, _ := usertools.RandomUser()
	url := dummyURL(t, user.Username)

	testCases := []struct {
		name        string
		username    string
		randomID    string
		buildStub   func(store *mockurlshortener.MockPostgreSQL)
		checkResult func(result string, err error)
	}{
		{
			name:     "OK",
			username: user.Username,
			buildStub: func(store *mockurlshortener.MockPostgreSQL) {
				store.EXPECT().
					CreateShortUrl(gomock.Any(), user.Username, url.LongUrl, url.ShortUrl).
					Times(1).
					Return(url, nil)
			},
			checkResult: func(result string, err error) {
				require.NoError(t, err)
				require.NotNil(t, result)

				require.Equal(t, result, url.LongUrl)
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

			usecase := NewUrlShortenerUsecase(store, config)
			response, err := usecase.CreateShortURL(context.TODO(), tc.username, tc.randomID)
			tc.checkResult(response, err)
		})
	}
}

func dummyURL(t *testing.T, owner string) (url model.Url) {
	randomID := tools.RandomAlphanumericString(5)

	url = model.Url{
		LongUrl:   "https://www.notion.so/stockbit/Backend-Engineering-Challenge-Link-Shortener-82bf71375701427c9cdd54a10a775ba6?pvs=4",
		ShortUrl:  randomID,
		CreatedAt: time.Now(),
		Owner:     owner,
	}

	return
}
