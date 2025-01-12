package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/internal/pkg/task"
	"VOU-Server/proto/gen"
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	DEFAULT_QUIZ_NUM   = 3
	DEFAULT_QUIZ_GENRE = "miscellaneous"

	QuizGameID  = 1
	ShakeGameID = 2
)

func (server *Server) CreateEvent(ctx context.Context, req *gen.CreateEventRequest) (*gen.CreateEventResponse, error) {
	arg := db.CreateEventTxParams{
		CreateEventParams: db.CreateEventParams{
			Owner:           req.GetOwner(),
			GameID:          req.GetGameId(),
			StoreID:         req.GetStoreId(),
			Name:            req.GetName(),
			VoucherQuantity: req.GetVoucherQuantity(),
			StartTime:       req.GetStartTime().AsTime(),
			EndTime:         req.GetEndTime().AsTime(),
		},
	}

	var (
		event db.Event
		err   error
	)
	if req.GetGameId() == ShakeGameID {
		arg.CreateEventParams.Status = db.EventsStatusNotAccepted
		event, err = server.store.CreateEvent(ctx, arg.CreateEventParams)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create event: %s", err)
		}
	} else if req.GetGameId() == QuizGameID {
		arg.CreateEventParams.Status = db.EventsStatusGenerating

		arg.AfterCreate = func(event db.Event) error {
			payload := task.PayloadGenQuiz{
				EventId: event.ID,
			}
			payload.QuizGenre = DEFAULT_QUIZ_GENRE
			if req.GetQuizGenre() != "" {
				payload.QuizGenre = req.GetQuizGenre()
			}

			payload.QuizNum = DEFAULT_QUIZ_NUM
			if req.GetQuizNumber() != 0 {
				payload.QuizNum = req.GetQuizNumber()
			}

			bytes, err := json.Marshal(payload)
			if err != nil {
				return fmt.Errorf("json.Marshal[event]: %w", err)
			}

			return server.publisher.Publish(ctx, bytes, "text/plain")
		}

		txResult, err := server.store.CreateEventTx(ctx, arg)
		if err != nil {
			// TODO: need more error handling
			return nil, status.Errorf(codes.Internal, "failed to create event: %s", err)
		}
		event = txResult.Event
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "invalid game id")
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
