package event

import (
	"VOU-Server/proto/gen"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServiceClient(eventSvcUrl string) (gen.EventServiceClient, error) {
	conn, err := grpc.NewClient(eventSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("cannot connect to event service: %w", err)
	}

	client := gen.NewEventServiceClient(conn)
	return client, nil
}
