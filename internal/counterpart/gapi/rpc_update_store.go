package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateStore(ctx context.Context, req *gen.UpdateStoreRequest) (*gen.UpdateStoreResponse, error) {
	store, err := server.store.UpdateStore(ctx, db.UpdateStoreParams{
		ID:    req.GetId(),
		Owner: req.GetOwner(),
		Name: pgtype.Text{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		BusinessType: pgtype.Text{
			String: req.GetBusinessType(),
			Valid:  req.BusinessType != nil,
		},
	})
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "store not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update store: %s", err)
	}

	rsp := &gen.UpdateStoreResponse{
		Store: &gen.Store{
			Id:           store.ID,
			Name:         store.Name,
			Owner:        store.Owner,
			BusinessType: store.BusinessType,
		},
	}
	return rsp, nil
}
