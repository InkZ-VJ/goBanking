DB_URL=postgres://postgres:postgres@localhost:5000/simple_bank?sslmode=disable

.PHONY: postgres go_app build-app redis
postgres: 
	docker compose up go_db
go_app:
	docker compose build && docker compose up go_app
build-app:
	docker compose --env-file ./app.env up -d
redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: createdb dropdb
createdb:
	docker exec -it go_db createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it go_db dropdb simple_bank

.PHONY: new_migration migrateUp migrateDown migrateUp1 migrateDown1 
new_migration:
	migrate create -ext sql -dir db/migration/ -seq ${name}
migrateUp:
	migrate -path db/migration -database "$(DB_URL)"  --verbose up
migrateDown:
	migrate -path db/migration -database "$(DB_URL)"  --verbose down
migrateUp1:
	migrate -path db/migration -database "$(DB_URL)"  --verbose up 1
migrateDown1:
	migrate -path db/migration -database "$(DB_URL)"  --verbose down 1

.PHONY:	sqlc test server mock proto evans
sqlc:
	sqlc generate
test:
	go test -v -cover -short ./...
fail_test:
	go test -v -cover -short ./... | grep FAIL
server:
	go run main.go
mock: 
	mockgen --package mockdb --destination db/mock/store.go github.com/VatJittiprasert/goBanking/db/sqlc Store
	mockgen --package mockwk --destination worker/mock/distributor.go github.com/VatJittiprasert/goBanking/worker TaskDistributor
proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--experimental_allow_proto3_optional \
	proto/*.proto
evans:
	evans --host localhost --port 9000 -r repl
rand_32key:
	openssl rand -hex 64 | head -c 32



