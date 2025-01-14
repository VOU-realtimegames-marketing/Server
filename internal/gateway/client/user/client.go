package user

import (
	"VOU-Server/proto/gen"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServiceClient(userSvcUrl string) (gen.UserServiceClient, error) {
	conn, err := grpc.NewClient(userSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("cannot connect to user service: %w", err)
	}

	client := gen.NewUserServiceClient(conn)
	return client, nil
}
