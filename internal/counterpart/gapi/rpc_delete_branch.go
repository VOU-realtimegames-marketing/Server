package gapi

import (
	"VOU-Server/proto/gen"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteBranch(ctx context.Context, req *gen.DeleteBranchRequest) (*gen.DeleteBranchResponse, error) {
	err := server.store.DeleteBranch(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete store: %s", err)
	}

	return nil, nil
}
