package main

import (
	"VOU-Server/cmd/quiz/config"
	"VOU-Server/internal/quiz/app"
	"VOU-Server/pkg/llms"
	"VOU-Server/pkg/rabbitmq"
	pkgConsumer "VOU-Server/pkg/rabbitmq/consumer"
	pkgPublisher "VOU-Server/pkg/rabbitmq/publisher"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	a, cleanup, err := app.InitApp(config.DBSource, rabbitmq.RabbitMQConnStr(config.RabbitMQAddress), llms.LLMApiKey(config.GeminiAPIKey))
	a.QuizCreatedPub.Configure(
		pkgPublisher.ExchangeName("quiz-created-exchange"),
		pkgPublisher.BindingKey("quiz-created-routing-key"),
		pkgPublisher.MessageTypeName("quiz-created"),
	)

	a.Consumer.Configure(
		pkgConsumer.ExchangeName("quiz-event-exchange"),
		pkgConsumer.QueueName("quiz-event-queue"),
		pkgConsumer.BindingKey("quiz-event-routing-key"),
		pkgConsumer.ConsumerTag("quiz-event-consumer"),
	)

	go func() {
		err := a.Consumer.StartConsumer(a.Worker)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to start Consumer")
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		cleanup()
		log.Info().Any("signal.Notify", v)
	case done := <-ctx.Done():
		cleanup()
		log.Info().Any("ctx.Done", done)
	}
}
