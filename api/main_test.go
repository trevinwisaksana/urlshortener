package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/trevinwisaksana/trevin-urlshortener/config"
	db "github.com/trevinwisaksana/trevin-urlshortener/db/sqlc"
	"github.com/trevinwisaksana/trevin-urlshortener/tools"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := config.Config{
		TokenSymmetricKey:   tools.RandomAlphanumericString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
