package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/trevinwisaksana/trevin-urlshortener/tools"
)

func createRandomUser(t *testing.T) User {
	hashPassword, err := tools.HashPassword(tools.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       tools.RandomUsername(),
		HashedPassword: hashPassword,
		FullName:       tools.RandomName(),
		Email:          tools.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	userCreated := createRandomUser(t)
	userRetrieved, err := testQueries.GetUser(context.Background(), userCreated.Username)

	require.NoError(t, err)
	require.NotEmpty(t, userRetrieved)

	require.Equal(t, userRetrieved.Username, userRetrieved.Username)
	require.Equal(t, userRetrieved.FullName, userRetrieved.FullName)
	require.Equal(t, userRetrieved.HashedPassword, userRetrieved.HashedPassword)
	require.Equal(t, userRetrieved.Email, userRetrieved.Email)
	require.WithinDuration(t, userRetrieved.CreatedAt, userRetrieved.CreatedAt, time.Second)
}
