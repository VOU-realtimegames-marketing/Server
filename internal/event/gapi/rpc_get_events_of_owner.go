package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertEvents(events []db.ListEventsOfOwnerRow) []*gen.Event {
	genEvents := make([]*gen.Event, len(events))
	for i, event := range events {
		genEvents[i] = &gen.Event{
			Id:              event.ID,
			Name:            event.Name,
			VoucherQuantity: event.VoucherQuantity,
			GameId:          event.GameID,
			StoreId:         event.StoreID,
			Status:          string(event.Status),
			StartTime:       timestamppb.New(event.StartTime),
			EndTime:         timestamppb.New(event.EndTime),
			GameType:        &event.GameType,
			Store:           &event.Store,
		}
	}
	return genEvents
}

func (server *Server) GetAllEventsOfOwner(ctx context.Context, req *gen.GetEventsOfOwnerRequest) (*gen.GetEventsOfOwnerResponse, error) {
	events, err := server.store.ListEventsOfOwner(ctx, req.GetOwner())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "no event found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get events: %s", err)
	}

	genEvents := convertEvents(events)

	rsp := &gen.GetEventsOfOwnerResponse{
		Events: genEvents,
	}

	return rsp, nil
}
