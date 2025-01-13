package gapi

import (
	"VOU-Server/proto/gen"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	counterpartRole = "partner"
	adminRole       = "admin"
)

// STORE API

func (server *Server) CreateStore(ctx context.Context, req *gen.CreateStoreRequest) (*gen.CreateStoreResponse, error) {
	res, err := server.AuthorizeUser(ctx, &gen.AuthorizeRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	if res.User.Role != counterpartRole {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	req.Owner = res.User.Username
	return server.counterpartClient.CreateStore(ctx, req)
}

func (server *Server) GetAllStoresOfOwner(ctx context.Context, req *gen.GetStoresOfOwnerRequest) (*gen.GetStoresOfOwnerResponse, error) {
	res, err := server.AuthorizeUser(ctx, &gen.AuthorizeRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	if res.User.Role != counterpartRole {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	req.Owner = res.User.Username
	return server.counterpartClient.GetAllStoresOfOwner(ctx, req)
}

func (server *Server) UpdateStore(ctx context.Context, req *gen.UpdateStoreRequest) (*gen.UpdateStoreResponse, error) {
	res, err := server.AuthorizeUser(ctx, &gen.AuthorizeRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	if res.User.Role != counterpartRole {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	req.Owner = res.User.Username
	return server.counterpartClient.UpdateStore(ctx, req)
}

func (server *Server) DeleteStore(ctx context.Context, req *gen.DeleteStoreRequest) (*gen.DeleteStoreResponse, error) {
	res, err := server.AuthorizeUser(ctx, &gen.AuthorizeRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	if res.User.Role != counterpartRole {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	req.Owner = res.User.Username
	return server.counterpartClient.DeleteStore(ctx, req)
}

// END STORE API

// BRANCH API

func (server *Server) CreateBranch(ctx context.Context, req *gen.CreateBranchRequest) (*gen.CreateBranchResponse, error) {
	res, err := server.AuthorizeUser(ctx, &gen.AuthorizeRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	if res.User.Role != counterpartRole {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	req.Owner = res.User.Username
	return server.counterpartClient.CreateBranch(ctx, req)
}

func (server *Server) GetBranchs(ctx context.Context, req *gen.GetBranchsRequest) (*gen.GetBranchsResponse, error) {
	res, err := server.AuthorizeUser(ctx, &gen.AuthorizeRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	if res.User.Role != counterpartRole {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	return server.counterpartClient.GetBranchs(ctx, req)
}

func (server *Server) DeleteBranch(ctx context.Context, req *gen.DeleteBranchRequest) (*gen.DeleteBranchResponse, error) {
	res, err := server.AuthorizeUser(ctx, &gen.AuthorizeRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	if res.User.Role != counterpartRole {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	return server.counterpartClient.DeleteBranch(ctx, req)
}

func (server *Server) GetCmsOverview(ctx context.Context, req *gen.GetCmsOverviewRequest) (*gen.GetCmsOverviewResponse, error) {
	res, err := server.AuthorizeUser(ctx, &gen.AuthorizeRequest{})

	log.Print("GetCmsOverview: ", res.User)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	if res.User.Role != counterpartRole {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized action: %s", err)
	}

	req.Owner = res.User.Username

	log.Print("GetCmsOverview_counterpart Owner: ", req.Owner)
	return server.counterpartClient.GetCmsOverview(ctx, req)

	// TODO: query real data from Database

	// Fake data giống với frontend
	// response := &gen.GetCmsOverviewResponse{
	// 	TotalStore:             10,
	// 	TotalBranch:            20,
	// 	TotalEvent:             123,
	// 	TotalUserPlay:          12513,
	// 	LastMonthTotalUserPlay: 13000,
	// }

	// return response, nil
}
