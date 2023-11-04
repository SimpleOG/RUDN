package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	db "rudnWebApp/db/sqlc"
	server2 "rudnWebApp/server"
	configs "rudnWebApp/util"
)

func main() {
	config, err := configs.InitConfig(".")
	if err != nil {
		log.Fatalf("cannot load config")
	}

	connPool, err := pgxpool.New(context.Background(), config.DBDSource)
	if err != nil {
		log.Fatalf("cannot connect to db")
	}

	store := db.NewStore(connPool)
	fmt.Println("Работает", store)
	server, err := server2.NewServer(config, store)
	if err != nil {
		log.Fatalln("cannot create server : ", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalln("cannot start server : ", err)
	}

}
