package main

import (
	"VOU-Server/cmd/event/config"
	"VOU-Server/internal/event/app"
	"VOU-Server/pkg/logger"
	"VOU-Server/pkg/rabbitmq"
	pkgConsumer "VOU-Server/pkg/rabbitmq/consumer"
	pkgPublisher "VOU-Server/pkg/rabbitmq/publisher"
	"net"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	grpcLogger := grpc.UnaryInterceptor(logger.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)

	prepareApp(config, grpcServer)

	// Start gRPC server
	listener, err := net.Listen("tcp", config.EventServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start the server")
	}
}

func prepareApp(config config.Config, grpcServer *grpc.Server) func() {
	a, cleanup, err := app.InitApp(config.DBSource, rabbitmq.RabbitMQConnStr(config.RabbitMQAddress), grpcServer)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init app")
	}

	a.Publisher.Configure(
		pkgPublisher.ExchangeName("quiz-event-exchange"),
		pkgPublisher.BindingKey("quiz-event-routing-key"),
		pkgPublisher.MessageTypeName("quiz-event-created"),
	)

	a.Consumer.Configure(
		pkgConsumer.ExchangeName("quiz-created-exchange"),
		pkgConsumer.QueueName("quiz-created-queue"),
		pkgConsumer.BindingKey("quiz-created-routing-key"),
		pkgConsumer.ConsumerTag("quiz-created-consumer"),
	)

	go func() {
		err1 := a.Consumer.StartConsumer(a.Worker)
		if err1 != nil {
			log.Fatal().Err(err1).Msg("failed to start Consumer")
		}
	}()

	return cleanup
}
