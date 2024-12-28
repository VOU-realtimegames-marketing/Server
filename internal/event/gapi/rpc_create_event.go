package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/internal/pkg/task"
	"VOU-Server/proto/gen"
	"context"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	QuizGameID  = 1
	ShakeGameID = 2
)

func (server *Server) CreateEvent(ctx context.Context, req *gen.CreateEventRequest) (*gen.CreateEventResponse, error) {
	arg := db.CreateEventParams{
		GameID:          req.GetGameId(),
		StoreID:         req.GetStoreId(),
		Name:            req.GetName(),
		VoucherQuantity: req.GetVoucherQuantity(),
		// StartTime:       req.GetStartTime().AsTime(),
		// EndTime:         req.GetEndTime().AsTime(),
	}

	if req.GetGameId() == QuizGameID {
		arg.Status = "generating"
	} else if req.GetGameId() == ShakeGameID {
		arg.Status = "not_accepted"
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "invalid game id")
	}

	event, err := server.store.CreateEvent(ctx, arg)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create event: %s", err)
	}

	// Publish quiz gen message
	if req.GetGameId() == QuizGameID {
		payload := task.PayloadGenQuiz{
			EventId: event.ID,
		}
		payload.QuizGenre = "miscellaneous"
		if req.GetQuizGenre() != "" {
			payload.QuizGenre = req.GetQuizGenre()
		}

		bytes, err := json.Marshal(payload)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "json.Marshal[event]: %s", err)
		}

		server.publisher.Publish(ctx, bytes, "text/plain")
	}

	rsp := &gen.CreateEventResponse{
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
