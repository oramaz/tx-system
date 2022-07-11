package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(
		"postgres", fmt.Sprintf("postgres://%s:%s@localhost:5432/tx_system?sslmode=disable",
			os.Getenv("POSTGRES_NAME"), os.Getenv("POSTGRES_PASSWORD")),
	)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
