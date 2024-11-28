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

.PHONY: new_migration sqlc
