package gapi

import (
	"VOU-Server/cmd/gateway/config"
	"VOU-Server/internal/gateway/client/auth"
	"VOU-Server/proto/gen"
	"context"
	"fmt"
)

type Server struct {
	gen.UnimplementedGatewayServer
	authClient gen.AuthServiceClient
}

func NewServer(config config.Config) (*Server, error) {
	authClient, err := auth.NewServiceClient(config.AuthServerAddress)
	if err != nil {
		return nil, fmt.Errorf("cannot create auth client: %w", err)
	}

	server := &Server{
		authClient: authClient,
	}
	return server, nil
}

func (server *Server) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	return server.authClient.CreateUser(ctx, req)
}
