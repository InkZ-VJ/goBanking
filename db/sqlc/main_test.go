package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/VatJittiprasert/goBanking/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot laod config:", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
