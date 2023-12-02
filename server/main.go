package main

import (
	"context"
	"fmt"
	db "rudnWebApp/db/sqlc"
	"rudnWebApp/util"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"log"
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
	now := time.Now()
	err = store.ReadItAll()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(time.Since(now))
	//serv, err := api.NewServer(config, store)
	//if err != nil {
	//	log.Fatalln("cannot create api : ", err)
	//}
	//err = serv.Start(config.ServerAddress)
	//if err != nil {
	//	log.Fatalln("cannot start api : ", err)
	//}

}
