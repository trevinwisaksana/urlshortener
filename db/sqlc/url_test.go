package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	tools "github.com/trevinwisaksana/trevin-urlshortener/tools"
)

func TestShortenURL(t *testing.T) {
	randomID := tools.RandomString(5)
	longUrl := "https://www.notion.so/stockbit/Backend-Engineering-Challenge-Link-Shortener-82bf71375701427c9cdd54a10a775ba6?pvs=4"
	baseUrl := "http://localhost:8080/"

	arg := CreateShortURLParams{
		ID:       randomID,
		LongUrl:  longUrl,
		ShortUrl: baseUrl + randomID,
	}

	result, err := testQueries.CreateShortURL(context.Background(), arg)

	require.NoError(t, err)
	require.NotNil(t, result)

	require.Equal(t, arg.ID, result.ID)
	require.Equal(t, arg.LongUrl, result.LongUrl)
	require.Equal(t, arg.ShortUrl, result.ShortUrl)
}
