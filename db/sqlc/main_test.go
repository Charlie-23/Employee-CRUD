package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pzanwar/employee/util"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:secret@localhost:5432/employee_data?sslmode=disable"
// )

var testStore Store

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../../")

	if err != nil {
		log.Fatal("Cannot load configs", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("Cannot connet to DB:", err)
	}

	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
