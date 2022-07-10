package main

import (
	"log"

	// "github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatal(fmt.Sprintf("Unable to connect to database: %v\n", err))
	// }
	// defer conn.Close(context.Background())

	// log.Fatal(api.GetServer(conn).Listen(":3000"))
}