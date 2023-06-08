package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/trevinwisaksana/trevin-urlshortener/config"
	db "github.com/trevinwisaksana/trevin-urlshortener/db/sqlc"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := config.Config{}

	server := NewServer(config, store)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
