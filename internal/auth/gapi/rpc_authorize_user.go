package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/metadata"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (server *Server) AuthorizeUser(ctx context.Context, req *gen.AuthorizeRequest) (*gen.AuthorizeResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	authHeader := values[0] // <authorization-type> <authorization-data>
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	// Log ownerId (payload.Username) ngay tại đây
	log.Printf("AuthorizeUser: ownerId=%s", payload.Username)

	user, err := server.store.GetUser(ctx, db.GetUserParams{
		Username: pgtype.Text{
			String: payload.Username,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("cannot get user: %s", err)
	}

	rsp := &gen.AuthorizeResponse{
		User: convertUser(user),
	}
	return rsp, nil
}
