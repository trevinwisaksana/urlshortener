package urlshortenerrepo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/trevinwisaksana/trevin-urlshortener/model"
	"github.com/trevinwisaksana/trevin-urlshortener/tools"
)

func generateShortURL(t *testing.T) (model.Url, error) {
	randomID := tools.RandomAlphanumericString(5)
	user := tools.RandomUsername()

	longUrl := "https://www.notion.so/stockbit/Backend-Engineering-Challenge-Link-Shortener-82bf71375701427c9cdd54a10a775ba6?pvs=4"

	arg := CreateShortURLParams{
		LongUrl:  longUrl,
		ShortUrl: randomID,
		Owner:    user,
	}

	result, err := testQueries.CreateShortURL(context.Background(), arg)
	require.NoError(t, err, randomID)
	require.NotNil(t, result)

	require.Equal(t, arg.LongUrl, result.LongUrl)
	require.Equal(t, arg.ShortUrl, result.ShortUrl)

	return result, err
}

func TestShortenURL(t *testing.T) {
	generateShortURL(t)
}

func TestGetLongURL(t *testing.T) {
	response, err := generateShortURL(t)
	require.NoError(t, err)

	result, err := testQueries.GetLongURL(context.Background(), response.ShortUrl)

	require.NoError(t, err)
	require.NotNil(t, result)

	require.Equal(t, response.LongUrl, result)
}

func TestUpdateShortURL(t *testing.T) {
	dummyData, err := generateShortURL(t)
	require.NoError(t, err)

	randomID := tools.RandomAlphanumericString(5)

	arg := UpdateShortURLParams{
		ShortUrl:        dummyData.ShortUrl,
		CustomShortlink: randomID,
	}

	result, err := testQueries.UpdateShortURL(context.Background(), arg)

	require.NoError(t, err)
	require.NotNil(t, result)

	require.Equal(t, dummyData.ID, result.ID)
	require.Equal(t, result.ShortUrl, arg.CustomShortlink)
	require.Equal(t, dummyData.LongUrl, result.LongUrl)
}

func TestDuplicateUpdateShortURL(t *testing.T) {
	dummyData, err := generateShortURL(t)
	require.NoError(t, err)

	secondDummyData, err := generateShortURL(t)
	require.NoError(t, err)

	arg := UpdateShortURLParams{
		ShortUrl:        dummyData.ShortUrl,
		CustomShortlink: secondDummyData.ShortUrl,
	}

	_, err = testQueries.UpdateShortURL(context.Background(), arg)
	require.Error(t, err)
}

func createRandomUser(t *testing.T) User {
	hashPassword, err := tools.HashPassword(tools.RandomAlphanumericString(6))
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
