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
	rm -f proto/gen/*/*.go
	protoc --proto_path=proto --go_out=proto/gen --go_opt=paths=source_relative \
    --go-grpc_out=proto/gen --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=allow_delete_body=true,paths=source_relative:proto/gen \
    proto/*/*.proto && ./flatten_gen.sh

wire:
	cd internal/quiz/app && wire && cd - && \
	cd internal/event/app && wire && cd -

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

rabbitmq:
	docker run -d --hostname vou-rabbit --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4.0-management-alpine

auth:
	cd cmd/auth && go run main.go

counterpart:
	cd cmd/counterpart && go run main.go

gateway:
	cd cmd/gateway && go run main.go

event:
	cd cmd/event && go run main.go

quiz:
	cd cmd/quiz && go run main.go

game:
	cd cmd/game && go run main.go

game_client:
	cd cmd/game_client && go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration sqlc proto wire redis rabbitmq auth counterpart gateway
