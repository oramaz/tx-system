include .env
export

DATABASE_URL=postgres://${POSTGRES_NAME}:${POSTGRES_PASSWORD}@localhost:5432/tx_system?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=${POSTGRES_NAME} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres:latest

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres tx_system

dropdb:
	docker exec -it postgres dropdb --username=postgres tx_system

migrateup: 
	migrate -path ./internal/db/migration -database "${DATABASE_URL}" -verbose up

migratedown: 
	migrate -path ./internal/db/migration -database "${DATABASE_URL}" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run cmd/main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server