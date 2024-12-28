//go:build wireinject
// +build wireinject

package app

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/internal/quiz/handler"
	"VOU-Server/pkg/llms"
	"VOU-Server/pkg/rabbitmq"
	pkgConsumer "VOU-Server/pkg/rabbitmq/consumer"
	pkgPublisher "VOU-Server/pkg/rabbitmq/publisher"
	"context"

	"github.com/google/generative-ai-go/genai"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

func InitApp(
	dbSource string,
	rabbitMQConnStr rabbitmq.RabbitMQConnStr,
	apiKey llms.LLMApiKey,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		storeDBFunc,
		rabbitMQFunc,
		pkgPublisher.EventPublisherSet,
		pkgConsumer.EventConsumerSet,
		genAIFunc,
		llms.GeminiSet,
		handler.QuizGenHandlerSet,
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

func genAIFunc(apiKey llms.LLMApiKey) (*genai.Client, func(), error) {
	client, err := llms.NewGenerativeAIClient(context.Background(), apiKey)
	if err != nil {
		return nil, nil, err
	}
	return client, func() { client.Close() }, nil
}
