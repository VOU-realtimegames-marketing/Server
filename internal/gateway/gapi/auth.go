package gapi

import (
	"VOU-Server/proto/gen"
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
	cookieHeader        = "grpcgateway-cookie"
)

func (server *Server) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	// server.authClient.Authorize(ctx, &gen.AuthorizeRequest{})
	return server.authClient.CreateUser(ctx, req)
}

func (server *Server) LoginUser(ctx context.Context, req *gen.LoginUserRequest) (*gen.LoginUserResponse, error) {
	return server.authClient.LoginUser(ctx, req)
}

func (server *Server) AuthorizeUser(ctx context.Context, req *gen.AuthorizeRequest) (*gen.AuthorizeResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	authHeaderValues := md.Get(authorizationHeader)
	cookieHeaderValues := md.Get(cookieHeader)

	authHeader := ""
	if len(authHeaderValues) != 0 {
		authHeader = authHeaderValues[0] // <authorization-type> <authorization-data>
	}
	if len(cookieHeaderValues) != 0 {
		authHeader = cookieHeaderValues[0] // <authorization-type> <authorization-data>
		fields := strings.Split(authHeader, "; ")
		accessTokenField := fields[0]
		accessToken := strings.Split(accessTokenField, "=")[1]
		authHeader = fmt.Sprintf("%s %s", authorizationBearer, accessToken)
	}
	if authHeader == "" {
		return nil, fmt.Errorf("missing authorization header")
	}
	ctx = metadata.AppendToOutgoingContext(ctx, authorizationHeader, authHeader)
	// server.authClient.Authorize(ctx, &gen.AuthorizeRequest{})
	return server.authClient.AuthorizeUser(ctx, req)
}

func (server *Server) RenewAccessToken(ctx context.Context, req *gen.RenewAccessTokenRequest) (*gen.RenewAccessTokenResponse, error) {
	return server.authClient.RenewAccessToken(ctx, req)
}

func (server *Server) VerifyEmail(ctx context.Context, req *gen.VerifyEmailRequest) (*gen.VerifyEmailResponse, error) {
	return server.authClient.VerifyEmail(ctx, req)
}
