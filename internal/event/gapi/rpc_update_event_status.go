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

func (server *Server) UpdateEventStatus(ctx context.Context, req *gen.UpdateEventStatusRequest) (*gen.UpdateEventStatusResponse, error) {
	event, err := server.store.UpdateEvent(ctx, db.UpdateEventParams{
		ID: req.GetId(),
		Status: db.NullEventsStatus{
			EventsStatus: db.EventsStatus(req.GetStatus()),
			Valid:        true,
		},
	})
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "event not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update event: %s", err)
	}

	rsp := &gen.UpdateEventStatusResponse{
		Event: &gen.Event{
			Id:              event.ID,
			Name:            event.Name,
			VoucherQuantity: event.VoucherQuantity,
			GameId:          event.GameID,
			StoreId:         event.StoreID,
			StartTime:       timestamppb.New(event.StartTime),
			EndTime:         timestamppb.New(event.EndTime),
			Status:          string(event.Status),
		},
	}
	return rsp, nil
}
