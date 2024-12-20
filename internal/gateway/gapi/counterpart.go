package gapi

import (
	"VOU-Server/proto/gen"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	counterpartRole = "partner"
)

func (server *Server) CreateStore(ctx context.Context, req *gen.CreateStoreRequest) (*gen.CreateStoreResponse, error) {
	res, err := server.AuthorizeUser(ctx, &gen.AuthorizeRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	if res.User.Role != counterpartRole {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	req.Owner = res.User.Username
	return server.counterpartClient.CreateStore(ctx, req)
}
