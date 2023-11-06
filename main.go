package main

import (
	"fmt"
	"rudnWebApp/util"
)

func main() {
	//config, err := configs.InitConfig(".")
	//if err != nil {
	//	log.Fatalf("cannot load config")
	//}
	//
	//connPool, err := pgxpool.New(context.Background(), config.DBDSource)
	//if err != nil {
	//	log.Fatalf("cannot connect to db")
	//}
	//
	//store := db.NewStore(connPool)
	//fmt.Println("Работает", store)
	//server, err := server2.NewServer(config, store)
	//if err != nil {
	//	log.Fatalln("cannot create server : ", err)
	//}
	//err = server.Start(config.ServerAddress)
	//if err != nil {
	//	log.Fatalln("cannot start server : ", err)
	//}
	result := make(chan int)

	for i := 0; i < 5; i++ {

		go func() {
			res := util.RandomNumber()
			result <- int(res)
		}()
	}
	defer close(result)
	for i := range result {
		fmt.Println(i)
	}

}
