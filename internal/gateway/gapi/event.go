package gapi

import (
	"VOU-Server/proto/gen"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateEvent(ctx context.Context, req *gen.CreateEventRequest) (*gen.CreateEventResponse, error) {
	res, err := server.AuthorizeUser(ctx, &gen.AuthorizeRequest{})
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorized action: %s", err)
	}

	if res.User.Role != counterpartRole {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorized action: %s", err)
	}

	return server.eventClient.CreateEvent(ctx, req)
}
