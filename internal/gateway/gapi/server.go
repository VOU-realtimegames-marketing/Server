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
	// server.authClient.Authorize(ctx, &gen.AuthorizeRequest{})
	return server.authClient.CreateUser(ctx, req)
}

func (server *Server) LoginUser(ctx context.Context, req *gen.LoginUserRequest) (*gen.LoginUserResponse, error) {
	return server.authClient.LoginUser(ctx, req)
}

func (server *Server) AuthorizeUser(ctx context.Context, req *gen.AuthorizeRequest) (*gen.AuthorizeResponse, error) {
	// server.authClient.Authorize(ctx, &gen.AuthorizeRequest{})
	return server.authClient.AuthorizeUser(ctx, req)
}
