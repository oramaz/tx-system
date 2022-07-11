package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/oramaz/tx-system/internal/api"
	db "github.com/oramaz/tx-system/internal/db/sqlc"
)

const (
	addr = "0.0.0.0:8080"
)

var dbAddr string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dbAddr = fmt.Sprintf("postgres://%s:%s@localhost:5432/tx_system?sslmode=disable",
		os.Getenv("POSTGRES_NAME"), os.Getenv("POSTGRES_PASSWORD"))

	conn, err := sql.Open("postgres", dbAddr)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(addr); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
