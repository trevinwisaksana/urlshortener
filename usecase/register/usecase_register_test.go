package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/trevinwisaksana/trevin-urlshortener/model"
	mockurlshortener "github.com/trevinwisaksana/trevin-urlshortener/repository/urlshortener/postgresql/mock"
	"github.com/trevinwisaksana/trevin-urlshortener/tools"
)

func TestRegisterAPI(t *testing.T) {
	user, _ := randomUser(t)

	testCases := []struct {
		name        string
		username    string
		fullName    string
		password    string
		email       string
		buildStub   func(store *mockurlshortener.MockPostgreSQL)
		checkResult func(user model.UserResponse, err error)
	}{
		{
			name:     "OK",
			username: user.Username,
			fullName: user.FullName,
			password: user.HashedPassword,
			email:    user.Email,
			buildStub: func(store *mockurlshortener.MockPostgreSQL) {
				store.EXPECT().
					CreateUser(gomock.Any(), user.Username, user.HashedPassword, user.FullName, user.Email).
					Times(1).
					Return(user, nil)
			},
			checkResult: func(response model.UserResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, user)

				require.Equal(t, response.Username, user.Username)
				require.Equal(t, response.FullName, user.FullName)
				require.Equal(t, response.Email, user.Email)
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

			usecase := NewUrlShortenerRegisterUsecase(store)
			response, err := usecase.Register(context.TODO(), tc.username, tc.password, tc.fullName, tc.email)
			tc.checkResult(response, err)
		})
	}
}

func randomUser(t *testing.T) (user model.User, password string) {
	password = tools.RandomAlphanumericString(6)
	hashedPassword, err := tools.HashPassword(password)
	require.NoError(t, err)

	user = model.User{
		Username:       tools.RandomUsername(),
		FullName:       tools.RandomUsername(),
		Email:          tools.RandomEmail(),
		HashedPassword: hashedPassword,
	}

	return
}
