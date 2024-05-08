package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"

	"github.com/pzanwar/employee/api"
	db "github.com/pzanwar/employee/db/sqlc"
	"github.com/pzanwar/employee/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load configs", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("Cannot connet to DB:", err)
	}

	store := db.NewStore(connPool)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("Cannot create server:", err)
	}
	err = server.Start(config.HTTPServerAddress)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}

}
