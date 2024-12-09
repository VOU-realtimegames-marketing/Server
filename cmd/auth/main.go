package main

import (
	"VOU-Server/cmd/auth/config"
	db "VOU-Server/db/sqlc"
	"VOU-Server/internal/auth/gapi"
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
	config, err := config.LoadConfig("./cmd/auth/config")
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
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create serve")
	}

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	gen.RegisterAuthServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.AuthServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create listene")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start a serve")
	}
}
