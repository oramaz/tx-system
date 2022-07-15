package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/oramaz/tx-system/internal/util"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../../.env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(
		"postgres", config.DatabaseURL,
	)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
