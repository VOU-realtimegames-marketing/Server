package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetQuizzesByEventId(ctx context.Context, req *gen.GetQuizzesByEventIdRequest) (*gen.GetQuizzesByEventIdResponse, error) {
	_, err := server.store.GetEventByIdAndOwner(ctx, db.GetEventByIdAndOwnerParams{
		ID:    req.GetEventId(),
		Owner: req.GetOwner(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot get quizzes of event that you do not own")
	}

	quiz, err := server.store.GetQuizzesByEventId(ctx, req.GetEventId())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "no quiz found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get quiz: %s", err)
	}

	quizContent := &structpb.ListValue{}
	err = quizContent.UnmarshalJSON([]byte(quiz.Content))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert to structpb: %s", err)
	}

	rsp := &gen.GetQuizzesByEventIdResponse{
		Quiz: &gen.Quiz{
			Id:        quiz.ID,
			EventId:   quiz.EventID,
			QuizGenre: quiz.QuizGenre,
			Content:   quizContent,
			CreatedAt: timestamppb.New(quiz.CreatedAt),
		},
	}

	return rsp, nil
}
