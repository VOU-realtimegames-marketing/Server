package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func convertUsers(users []db.User) []*gen.User {
	genUsers := make([]*gen.User, len(users))
	for i, user := range users {
		genUsers[i] = &gen.User{
			Username: user.Username,
			FullName: user.FullName,
			Role:     user.Role,
		}
	}
	return genUsers
}

func (server *Server) GetAllOtherUsers(ctx context.Context, req *gen.GetAllOtherUsersRequest) (*gen.GetAllOtherUsersResponse, error) {
	users, err := server.store.ListAllOtherUsers(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "no user found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get other users: %s", err)
	}

	genUsers := convertUsers(users)
	rsp := &gen.GetAllOtherUsersResponse{
		Users: genUsers,
	}

	return rsp, nil
}
