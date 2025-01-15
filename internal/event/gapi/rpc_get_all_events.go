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

func convertAllEvents(events []db.ListEventsRow) []*gen.Event {
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
			Store:           &event.Store,
			GameType:        &event.GameType,
			QuizNum:         &event.QuizNum.Int32,
		}
	}
	return genEvents
}

func (server *Server) GetAllEvents(ctx context.Context, req *gen.GetAllEventsRequest) (*gen.GetAllEventsResponse, error) {
	events, err := server.store.ListEvents(ctx)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "no event found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get events: %s", err)
	}

	genEvents := convertAllEvents(events)

	rsp := &gen.GetAllEventsResponse{
		Events: genEvents,
	}

	return rsp, nil
}
