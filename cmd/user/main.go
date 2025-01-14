package main

import (
	"VOU-Server/cmd/user/config"
	db "VOU-Server/db/sqlc"
	"VOU-Server/internal/user/gapi"
	"VOU-Server/pkg/logger"
	"VOU-Server/proto/gen"
	"context"
	"net"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to DB")
	}

	store := db.NewStore(connPool)
	runGrpcServer(config, store)
}

func runGrpcServer(config config.Config, store db.StoreDB) {
	server, err := gapi.NewServer(store)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create server")
	}

	grpcLogger := grpc.UnaryInterceptor(logger.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	gen.RegisterUserServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.UserServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start the server")
	}
}
