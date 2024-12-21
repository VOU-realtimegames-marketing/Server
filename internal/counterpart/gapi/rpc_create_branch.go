package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBranch(ctx context.Context, req *gen.CreateBranchRequest) (*gen.CreateBranchResponse, error) {
	_, err := server.store.GetStoreByIdAndOwner(ctx, db.GetStoreByIdAndOwnerParams{
		ID:    req.GetStoreId(),
		Owner: req.GetOwner(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot create branch that you do not own")
	}

	branch, err := server.store.CreateBranch(ctx, db.CreateBranchParams{
		StoreID:  req.GetStoreId(),
		Name:     req.GetName(),
		Position: req.GetPosition(),
		CityName: req.GetCityName(),
		Country:  req.GetCountry(),
		Address:  req.GetAddress(),
		Emoji:    req.GetEmoji(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create branch: %s", err)
	}

	rsp := &gen.CreateBranchResponse{
		Branch: &gen.Branch{
			Id:       branch.ID,
			StoreId:  branch.StoreID,
			Name:     branch.Name,
			Position: branch.Position,
			CityName: branch.CityName,
			Country:  branch.Country,
			Address:  branch.Address,
			Emoji:    branch.Emoji,
		},
	}
	return rsp, nil
}
