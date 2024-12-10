package auth

import (
	"VOU-Server/proto/gen"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServiceClient(authSvcUrl string) (gen.AuthServiceClient, error) {
	conn, err := grpc.NewClient(authSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("cannot connect to auth service: %w", err)
	}

	client := gen.NewAuthServiceClient(conn)
	return client, nil
}
