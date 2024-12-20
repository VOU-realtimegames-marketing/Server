package counterpart

import (
	"VOU-Server/proto/gen"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServiceClient(counterpartSvcUrl string) (gen.CounterpartServiceClient, error) {
	conn, err := grpc.NewClient(counterpartSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("cannot connect to counterpart service: %w", err)
	}

	client := gen.NewCounterpartServiceClient(conn)
	return client, nil
}
