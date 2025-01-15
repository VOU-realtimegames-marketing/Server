package gapi

import (
	"VOU-Server/proto/gen"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetAllOtherUsers(ctx context.Context, req *gen.GetAllOtherUsersRequest) (*gen.GetAllOtherUsersResponse, error) {
	res, err := server.AuthorizeUser(ctx, &gen.AuthorizeRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	req.Username = res.User.Username
	return server.userClient.GetAllOtherUsers(ctx, req)
}
