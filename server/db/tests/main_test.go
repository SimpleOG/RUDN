package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"rudnWebApp/db/sqlc"
	configs "rudnWebApp/util"
	"sync"

	"log"
	"testing"
)

var testQueries db.Store
var testDB *pgxpool.Pool

func gorutineCreate(fn func()) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fn()
			wg.Done()
		}()
	}
	wg.Wait()
}
func TestMain(m *testing.M) {
	var err error
	config, err := configs.InitConfig(".")
	if err != nil {
		log.Fatalln("cannot load config :", err)
	}
	testDB, err = pgxpool.New(context.Background(), config.DBDSource)
	if err != nil {
		log.Fatalln("cannot connect to db :", err)
	}
	testQueries = db.NewStore(testDB)
	os.Exit(m.Run())
}
