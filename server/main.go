package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"log"
	"rudnWebApp/api"
	db "rudnWebApp/db/sqlc"
	"rudnWebApp/util"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config, err := util.InitConfig(".")
	if err != nil {
		log.Fatalf("cannot load config %s", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBDSource)
	if err != nil {
		log.Fatalf("cannot connect to db %s", err)
	}
	store := db.NewStore(connPool)
	runDBMigration(config.MigrationUrl, config.DBDSource)
	serv, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalln("cannot create api : ", err)
	}

	err = serv.Start(config.ServerAddress)
	if err != nil {
		log.Fatalln("cannot start api : ", err)
	}

}

func runDBMigration(migrationURL, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatalln("cannot create migration", err)
	}
	if err = migration.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("уже заполнено")
			return
		}
		log.Fatalln("cannot start migration", err)
	}

}
