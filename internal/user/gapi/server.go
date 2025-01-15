package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
)

type Server struct {
	gen.UnimplementedUserServiceServer
	store db.StoreDB
}

// NewServer creates a new User gRPC server
func NewServer(store db.StoreDB) (*Server, error) {
	server := &Server{
		store: store,
	}
	return server, nil
}
