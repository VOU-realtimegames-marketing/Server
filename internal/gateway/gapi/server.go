package gapi

import (
	"VOU-Server/cmd/gateway/config"
	"VOU-Server/internal/gateway/client/auth"
	"VOU-Server/internal/gateway/client/counterpart"
	"VOU-Server/proto/gen"
	"fmt"
)

type Server struct {
	gen.UnimplementedGatewayServer
	authClient        gen.AuthServiceClient
	counterpartClient gen.CounterpartServiceClient
}

func NewServer(config config.Config) (*Server, error) {
	authClient, err := auth.NewServiceClient(config.AuthServerAddress)
	counterpartClient, err := counterpart.NewServiceClient(config.CounterpartServerAddress)
	if err != nil {
		return nil, fmt.Errorf("cannot create auth client: %w", err)
	}

	server := &Server{
		authClient:        authClient,
		counterpartClient: counterpartClient,
	}
	return server, nil
}
