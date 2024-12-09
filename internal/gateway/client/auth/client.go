package auth

import (
	"VOU-Server/proto/gen"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	Client gen.AuthServiceClient
}

func NewServiceClient(authSvcUrl string) *AuthServiceClient {
	conn, err := grpc.NewClient(authSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &AuthServiceClient{
		Client: gen.NewAuthServiceClient(conn),
	}
}
