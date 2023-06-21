package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/trevinwisaksana/trevin-urlshortener/config"
	db "github.com/trevinwisaksana/trevin-urlshortener/db/sqlc"
	"github.com/trevinwisaksana/trevin-urlshortener/token"
)

type Server struct {
	config     config.Config
	tokenMaker token.PasetoMaker
	store      db.Store
	router     *gin.Engine
}

func NewServer(config config.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: *tokenMaker,
		store:      store,
	}

	router := gin.Default()
	authRouter := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRouter.POST("/shortener", server.createShortURL)
	authRouter.POST("/edit", server.editURL)
	router.GET("/:id", server.redirectURL)

	router.POST("/register", server.createUser)
	router.POST("/login", server.login)

	server.router = router
	return server, err
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
