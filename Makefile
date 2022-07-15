include .env
export

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

mock:
	mockgen -package mockdb -destination ./internal/db/mock/store.go github.com/oramaz/tx-system/internal/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock