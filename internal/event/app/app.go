package app

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/internal/event/gapi"
	"VOU-Server/internal/event/handler"
	"VOU-Server/internal/pkg/task"
	pkgConsumer "VOU-Server/pkg/rabbitmq/consumer"
	pkgPublisher "VOU-Server/pkg/rabbitmq/publisher"
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type App struct {
	AMQPConn *amqp.Connection
	store    db.StoreDB

	Publisher pkgPublisher.EventPublisher
	Consumer  pkgConsumer.EventConsumer

	handler handler.QuizCreatedHandler
	server  *gapi.Server
}

func New(
	amqpConn *amqp.Connection,
	store db.StoreDB,
	publisher pkgPublisher.EventPublisher,
	consumer pkgConsumer.EventConsumer,
	handler handler.QuizCreatedHandler,
	server *gapi.Server,
) *App {
	return &App{
		AMQPConn: amqpConn,
		store:    store,

		Publisher: publisher,
		Consumer:  consumer,
		handler:   handler,
		server:    server,
	}
}

func (a *App) Worker(ctx context.Context, messages <-chan amqp.Delivery) {
	for delivery := range messages {
		log.Info().Uint64("delivery_tag", delivery.DeliveryTag).Msg("processDeliveries")
		log.Info().Str("delivery_type", delivery.Type).Msg("received")

		switch delivery.Type {
		case "quiz-created":
			var payload task.PayloadQuizCreated

			err := json.Unmarshal(delivery.Body, &payload)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to Unmarshal")
			}

			err = a.handler.Handle(ctx, payload)

			if err != nil {
				if err = delivery.Reject(false); err != nil {
					log.Fatal().Err(err).Msg("failed to delivery.Reject")
				}

				log.Fatal().Err(err).Msg("failed to process delivery")
			} else {
				err = delivery.Ack(false)
				if err != nil {
					log.Fatal().Err(err).Msg("failed to acknowledge delivery")

				}
			}

		default:
			log.Info().Msg("default")
		}
	}

	log.Info().Msg("Deliveries channel closed")
}
