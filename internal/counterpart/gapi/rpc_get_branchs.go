package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func convertBranchs(branchs []db.Branch) []*gen.Branch {
	genBranchs := make([]*gen.Branch, len(branchs))
	for i, branch := range branchs {
		genBranchs[i] = &gen.Branch{
			Id:       branch.ID,
			StoreId:  branch.StoreID,
			Name:     branch.Name,
			Position: branch.Position,
			CityName: branch.CityName,
			Country:  branch.Country,
			Address:  branch.Address,
			Emoji:    branch.Emoji,
		}
	}
	return genBranchs
}

func (server *Server) GetBranchs(ctx context.Context, req *gen.GetBranchsRequest) (*gen.GetBranchsResponse, error) {
	branchs, err := server.store.ListBranchs(ctx, req.GetStoreId())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "no branch found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get branchs: %s", err)
	}

	genBranchs := convertBranchs(branchs)
	rsp := &gen.GetBranchsResponse{
		Branchs: genBranchs,
	}

	return rsp, nil
}
