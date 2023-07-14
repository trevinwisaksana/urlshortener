package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockurlshortener "github.com/trevinwisaksana/trevin-urlshortener/repository/urlshortener/postgresql/mock"
	urltools "github.com/trevinwisaksana/trevin-urlshortener/tools/url"
	usertools "github.com/trevinwisaksana/trevin-urlshortener/tools/user"
)

func TestRedirectURL(t *testing.T) {
	username := usertools.RandomUsername()
	url := urltools.DummyURL(username)

	testCases := []struct {
		name        string
		shortlink   string
		buildStub   func(store *mockurlshortener.MockPostgreSQL)
		checkResult func(result string, err error)
	}{
		{
			name:      "OK",
			shortlink: url.ShortUrl,
			buildStub: func(store *mockurlshortener.MockPostgreSQL) {
				store.EXPECT().
					GetURL(gomock.Any(), gomock.Eq(url.ShortUrl)).
					Times(1)
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

			usecase := NewUrlShortenerRedirectUsecase(store)
			response, err := usecase.GetURL(context.TODO(), tc.shortlink)
			tc.checkResult(response.LongUrl, err)
		})
	}
}
