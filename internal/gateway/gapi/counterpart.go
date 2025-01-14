package gapi

import (
	"VOU-Server/proto/gen"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	counterpartRole = "partner"
	adminRole       = "admin"
	userRole        = "user"
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
