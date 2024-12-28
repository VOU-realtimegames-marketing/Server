package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/pkg/rabbitmq/publisher"
	"VOU-Server/proto/gen"

	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	gen.UnimplementedEventServiceServer
	store     db.StoreDB
	publisher publisher.EventPublisher
}

var EventGRPCServerSet = wire.NewSet(NewServer)

// NewServer creates a new Auth gRPC server
func NewServer(grpcServer *grpc.Server, store db.StoreDB, publisher publisher.EventPublisher) *Server {
	server := &Server{
		store:     store,
		publisher: publisher,
	}

	gen.RegisterEventServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
	return server
}
