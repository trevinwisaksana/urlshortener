package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	tools "github.com/trevinwisaksana/trevin-urlshortener/tools"
)

func generateShortURL(t *testing.T) (Url, error) {
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
