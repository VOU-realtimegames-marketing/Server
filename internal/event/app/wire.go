//go:build wireinject
// +build wireinject

package app

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/internal/event/gapi"
	"VOU-Server/internal/event/handler"
	"VOU-Server/pkg/rabbitmq"
	pkgConsumer "VOU-Server/pkg/rabbitmq/consumer"
	pkgPublisher "VOU-Server/pkg/rabbitmq/publisher"
	"context"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func InitApp(
	dbSource string,
	rabbitMQConnStr rabbitmq.RabbitMQConnStr,
	grpcServer *grpc.Server,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		storeDBFunc,
		rabbitMQFunc,
		pkgPublisher.EventPublisherSet,
		pkgConsumer.EventConsumerSet,

		gapi.EventGRPCServerSet,
		handler.QuizCreatedHandlerSet,
	))
}

func storeDBFunc(dbSource string) db.StoreDB {
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to DB")
	}

	store := db.NewStore(connPool)
	return store
}

func rabbitMQFunc(url rabbitmq.RabbitMQConnStr) (*amqp.Connection, func(), error) {
	conn, err := rabbitmq.NewRabbitMQConn(url)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}
