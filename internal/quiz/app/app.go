package app

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/internal/pkg/task"
	"VOU-Server/internal/quiz/handler"
	pkgConsumer "VOU-Server/pkg/rabbitmq/consumer"
	pkgPublisher "VOU-Server/pkg/rabbitmq/publisher"
	"context"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/genai"
)

type App struct {
	AMQPConn    *amqp.Connection
	store       db.StoreDB
	genAIClient *genai.Client

	QuizCreatedPub pkgPublisher.EventPublisher
	Consumer       pkgConsumer.EventConsumer

	handler handler.QuizGenHandler
}

func New(
	amqpConn *amqp.Connection,
	store db.StoreDB,
	genAIClient *genai.Client,
	quizCreatedPub pkgPublisher.EventPublisher,
	consumer pkgConsumer.EventConsumer,
	handler handler.QuizGenHandler,
) *App {
	return &App{
		AMQPConn:    amqpConn,
		store:       store,
		genAIClient: genAIClient,

		QuizCreatedPub: quizCreatedPub,
		Consumer:       consumer,

		handler: handler,
	}
}

func (c *App) Worker(ctx context.Context, messages <-chan amqp091.Delivery) {
	for delivery := range messages {
		log.Info().Uint64("delivery_tag", delivery.DeliveryTag).Msg("processDeliveries")
		log.Info().Str("delivery_type", delivery.Type).Msg("received")

		switch delivery.Type {
		case "quiz-event-generated":
			var payload task.PayloadGenQuiz

			err := json.Unmarshal(delivery.Body, &payload)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to Unmarshal")
			}

			err = c.handler.Handle(ctx, payload)

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
