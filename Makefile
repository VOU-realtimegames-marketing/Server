DB_URL=postgresql://root:secret@localhost:5432/vou-marketing?sslmode=disable

network:
	docker network create vou-network

postgres:
	docker run --name pgvou --network vou-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it pgvou createdb --username=root --owner=root vou-marketing

dropdb:
	docker exec -it pgvou dropdb vou-marketing

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

proto:
	rm -f proto/gen/*.go
	protoc --proto_path=proto --go_out=proto/gen --go_opt=paths=source_relative \
    --go-grpc_out=proto/gen --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=proto/gen --grpc-gateway_opt=paths=source_relative \
    proto/*.proto

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration sqlc proto
