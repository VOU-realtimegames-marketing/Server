package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
)

type Server struct {
	gen.UnimplementedGameServiceServer
	store db.StoreDB
}

// NewServer creates a new Game gRPC server
func NewServer(store db.StoreDB) (*Server, error) {
	server := &Server{
		store: store,
	}
	return server, nil
}
