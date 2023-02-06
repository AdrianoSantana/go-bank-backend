postgres:
	docker run --name bank-go-db -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -p 5432:5432 -d postgres:12-alpine
createdb:
	docker exec -it bank-go-db createdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

dropdb:
	docker exec -it bank-go-db dropdb simple_bank sqlc

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb