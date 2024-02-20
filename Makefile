DB_URL=postgres://postgres:postgres@localhost:5000/simple_bank?sslmode=disable

postgres: 
	docker compose up go_db
go_app:
	docker compose build && docker compose up go_app

build-app:
	docker compose --env-file ./app.env up -d

createdb:
	docker exec -it go_db createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it go_db dropdb simple_bank

initMigration:
	migrate create -ext sql -dir db/migration/ -seq ${INIT}
migrateUp:
	migrate -path db/migration -database "$(DB_URL)"  --verbose up
migrateDown:
	migrate -path db/migration -database "$(DB_URL)"  --verbose down
migrateUp1:
	migrate -path db/migration -database "$(DB_URL)"  --verbose up 1
migrateDown1:
	migrate -path db/migration -database "$(DB_URL)"  --verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen --package mockdb --destination db/mock/store.go github.com/VatJittiprasert/goBanking/db/sqlc Store

.PHONY:	postgres go_app build-app createdb dropdb migrateUp migrateDown migrateInit sqlc server mock migrateUp1 migrateDown1
