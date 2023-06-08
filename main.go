package main

import (
	"database/sql"
	"log"

	"github.com/trevinwisaksana/trevin-urlshortener/api"
	"github.com/trevinwisaksana/trevin-urlshortener/config"
	db "github.com/trevinwisaksana/trevin-urlshortener/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(config, store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
