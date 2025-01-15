package gapi

import (
	"VOU-Server/cmd/gateway/config"
	"VOU-Server/internal/gateway/client/auth"
	"VOU-Server/internal/gateway/client/counterpart"
	"VOU-Server/internal/gateway/client/event"
	"VOU-Server/internal/gateway/client/user"
	"VOU-Server/proto/gen"
	"fmt"
)

type Server struct {
	gen.UnimplementedGatewayServer
	authClient        gen.AuthServiceClient
	counterpartClient gen.CounterpartServiceClient
	eventClient       gen.EventServiceClient
	userClient        gen.UserServiceClient
}

func NewServer(config config.Config) (*Server, error) {
	authClient, err := auth.NewServiceClient(config.AuthServerAddress)
	counterpartClient, err := counterpart.NewServiceClient(config.CounterpartServerAddress)
	eventClient, err := event.NewServiceClient(config.EventServerAddress)
	userClient, err := user.NewServiceClient(config.UserServerAddress)

	if err != nil {
		return nil, fmt.Errorf("cannot create auth client: %w", err)
	}

	server := &Server{
		authClient:        authClient,
		counterpartClient: counterpartClient,
		eventClient:       eventClient,
		userClient:        userClient,
	}
	return server, nil
}
