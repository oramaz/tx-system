ifneq (,$(wildcard ./.env))
    include .env
    export
endif

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=${POSTGRES_NAME} -e POSTGRES_PASSWORD=${POSTGRES_PASSWD} -d postgres:latest

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

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test