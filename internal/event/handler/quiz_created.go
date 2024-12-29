package handler

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/internal/pkg/task"
	"context"

	"github.com/google/wire"
	"github.com/rs/zerolog/log"
)

var _ QuizCreatedHandler = (*quizCreatedHandler)(nil)

var QuizCreatedHandlerSet = wire.NewSet(NewQuizCreatedHandler)

type quizCreatedHandler struct {
	store db.StoreDB
}

func NewQuizCreatedHandler(store db.StoreDB) QuizCreatedHandler {
	return &quizCreatedHandler{
		store: store,
	}
}

func (h *quizCreatedHandler) Handle(ctx context.Context, payload task.PayloadQuizCreated) error {
	log.Info().Int64("EventID: ", payload.EventId).Msg("received event: quiz created")

	// update event status to 'not_accepted'
	_, err := h.store.UpdateEvent(ctx, db.UpdateEventParams{
		ID: payload.EventId,
		Status: db.NullEventsStatus{
			EventsStatus: db.EventsStatusNotAccepted,
			Valid:        true,
		},
	})

	return err
}
