package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"rudnWebApp/api"
	db "rudnWebApp/db/sqlc"
	"rudnWebApp/util"
)

func main() {
	config, err := util.InitConfig(".")
	if err != nil {
		log.Fatalf("cannot load config")
	}
	connPool, err := pgxpool.New(context.Background(), config.DBDSource)
	if err != nil {
		log.Fatalf("cannot connect to db")
	}
	store := db.NewStore(connPool)

	serv, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalln("cannot create api : ", err)
	}
	err = serv.Start(config.ServerAddress)
	if err != nil {
		log.Fatalln("cannot start api : ", err)
	}

}
