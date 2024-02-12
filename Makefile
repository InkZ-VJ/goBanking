postgres: 
	docker-compose up go_db

createdb:
	docker exec -it go_db createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it go_db dropdb simple_bank

initMigration:
	migrate create -ext sql -dir db/migration/ -seq ${INIT}
migrateUp:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5000/simple_bank?sslmode=disable" --verbose up
migrateDown:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5000/simple_bank?sslmode=disable" --verbose down
migrateUp1:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5000/simple_bank?sslmode=disable" --verbose up 1
migrateDown1:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5000/simple_bank?sslmode=disable" --verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen --package mockdb --destination db/mock/store.go github.com/VatJittiprasert/goBanking/db/sqlc Store

.PHONY:	postgres createdb dropdb migrateUp migrateDown migrateInit sqlc server mock migrateUp1 migrateDown1
