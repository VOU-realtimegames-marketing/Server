package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteStore(ctx context.Context, req *gen.DeleteStoreRequest) (*gen.DeleteStoreResponse, error) {
	err := server.store.DeleteStore(ctx, db.DeleteStoreParams{
		ID:    req.GetId(),
		Owner: req.GetOwner(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete store: %s", err)
	}

	return nil, nil
}
