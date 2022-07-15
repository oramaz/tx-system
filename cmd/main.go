package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/oramaz/tx-system/internal/api"
	db "github.com/oramaz/tx-system/internal/db/sqlc"
	"github.com/oramaz/tx-system/internal/util"
)

func main() {
	config, err := util.LoadConfig(".env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
