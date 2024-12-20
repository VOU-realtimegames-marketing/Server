package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func convertStores(stores []db.Store) []*gen.Store {
	genStores := make([]*gen.Store, len(stores))
	for i, store := range stores {
		genStores[i] = &gen.Store{
			Id:           store.ID,
			Name:         store.Name,
			Owner:        store.Owner,
			BusinessType: store.BusinessType,
		}
	}
	return genStores
}

func (server *Server) GetAllStoresOfOwner(ctx context.Context, req *gen.GetStoresOfOwnerRequest) (*gen.GetStoresOfOwnerResponse, error) {
	stores, err := server.store.ListStoresOfOwner(ctx, req.GetOwner())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "no store found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get stores: %s", err)
	}

	genStores := convertStores(stores)

	rsp := &gen.GetStoresOfOwnerResponse{
		Stores: genStores,
	}

	return rsp, nil
}
