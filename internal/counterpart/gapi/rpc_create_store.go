package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateStore(ctx context.Context, req *gen.CreateStoreRequest) (*gen.CreateStoreResponse, error) {
	store, err := server.store.CreateStore(ctx, db.CreateStoreParams{
		Owner:        req.GetOwner(),
		Name:         req.GetName(),
		BusinessType: req.GetBusinessType(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create store: %s", err)
	}

	rsp := &gen.CreateStoreResponse{
		Store: &gen.Store{
			Id:           store.ID,
			Name:         store.Name,
			Owner:        store.Owner,
			BusinessType: store.BusinessType,
		},
	}
	return rsp, nil
}
