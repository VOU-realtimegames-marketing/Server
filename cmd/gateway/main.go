package main

import (
	"VOU-Server/cmd/gateway/config"
	"VOU-Server/internal/gateway/client/auth"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := config.LoadConfig("./cmd/gateway/config")
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	router := gin.Default()

	authSvc := auth.NewServiceClient(config.AuthServerAddress)
	authSvc.InitRoutes(router, config)
}
