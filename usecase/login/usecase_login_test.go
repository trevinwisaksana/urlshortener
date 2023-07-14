package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/trevinwisaksana/trevin-urlshortener/config"
	mockurlshortener "github.com/trevinwisaksana/trevin-urlshortener/repository/urlshortener/postgresql/mock"
	usertools "github.com/trevinwisaksana/trevin-urlshortener/tools/user"
)

func TestLoginAPI(t *testing.T) {
	user, _ := usertools.RandomUser()

	testCases := []struct {
		name        string
		username    string
		password    string
		buildStub   func(store *mockurlshortener.MockPostgreSQL)
		checkResult func(response LoginUserResponse, err error)
	}{
		{
			name: "OK",
			buildStub: func(store *mockurlshortener.MockPostgreSQL) {
				store.EXPECT().
					CreateUser(gomock.Any(), user.Username, user.HashedPassword, user.FullName, user.Email).
					Times(1).
					Return(user, nil)
			},
			checkResult: func(response LoginUserResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, response)
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
			tokenMaker := tokenMaker.CreateToken(user.Username, config.AccessTokenDuration)

			usecase := NewUrlShortenerLoginUsecase(store, tokenMaker, config)
			response, err := usecase.Login(context.TODO(), tc.username, tc.password)
			tc.checkResult(response, err)
		})
	}
}
