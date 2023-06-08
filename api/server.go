package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trevinwisaksana/trevin-urlshortener/config"
	db "github.com/trevinwisaksana/trevin-urlshortener/db/sqlc"
)

type Server struct {
	config config.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config config.Config, store db.Store) *Server {
	server := &Server{
		config: config,
		store:  store,
	}

	router := gin.Default()

	router.POST("/shortener", server.createShortURL)
	router.GET("/:id", server.redirectURL)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
