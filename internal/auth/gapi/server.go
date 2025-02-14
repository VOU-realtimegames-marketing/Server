package gapi

import (
	"VOU-Server/cmd/auth/config"
	db "VOU-Server/db/sqlc"
	"VOU-Server/internal/auth/token"
	"VOU-Server/internal/auth/worker"
	"VOU-Server/proto/gen"
	"fmt"
)

type Server struct {
	gen.UnimplementedAuthServiceServer
	config          config.Config
	tokenMaker      token.Maker
	store           db.StoreDB
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new Auth gRPC server
func NewServer(config config.Config, store db.StoreDB, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		tokenMaker:      tokenMaker,
		store:           store,
		taskDistributor: taskDistributor,
	}
	return server, nil
}
