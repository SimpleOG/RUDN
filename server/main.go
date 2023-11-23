package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"rudnWebApp/server/api"
	db "rudnWebApp/server/db/sqlc"
	configs "rudnWebApp/server/util"
)

func main() {
	config, err := configs.InitConfig("C:\\Users\\Oleg\\GolandProjects\\rudnWebApp\\server/")
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
	//now := time.Now()
	//err = store.MakeItAll()
	//if err != nil {
	//	log.Fatalln(" пизда : ", err)
	//}
	//dur := time.Since(now)
	//fmt.Println(dur)

}
