package handler

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/internal/pkg/task"
	"VOU-Server/pkg/llms"
	"VOU-Server/pkg/rabbitmq/publisher"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var _ QuizGenHandler = (*quizGenHandler)(nil)

var QuizGenHandlerSet = wire.NewSet(NewQuizGenHandler)

type quizGenHandler struct {
	eventPub publisher.EventPublisher
	store    db.StoreDB
	gemini   llms.GenAI
}

func NewQuizGenHandler(eventPub publisher.EventPublisher, store db.StoreDB, gemini llms.GenAI) QuizGenHandler {
	return &quizGenHandler{
		eventPub: eventPub,
		store:    store,
		gemini:   gemini,
	}
}

func (h *quizGenHandler) Handle(ctx context.Context, payload task.PayloadGenQuiz) error {
	log.Info().Msg("received event: quiz event generated")

	fmt.Println("payload: ", payload)

	quizzesStr, err := h.gemini.GenerateContent(
		ctx,
		fmt.Sprintf(task.QuizGeneratePrompt, payload.QuizNum, payload.QuizGenre),
	)
	jsonStr := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(quizzesStr, "```json"), "```\n"))
	fmt.Println("jsonStr: ", jsonStr)

	var quizzes []task.Quiz
	err = json.Unmarshal([]byte(jsonStr), &quizzes)
	fmt.Println("quizzes: ", quizzes)

	arg := db.CreateQuizTxParams{
		CreateQuizParams: db.CreateQuizParams{
			EventID:   payload.EventId,
			QuizGenre: payload.QuizGenre,
			Content:   []byte(jsonStr),
		},
		AfterCreate: func(quiz db.Quiz) error {
			payloadQuizCreated := task.PayloadQuizCreated{
				EventId: quiz.EventID,
			}
			bytes, err := json.Marshal(payloadQuizCreated)
			if err != nil {
				return errors.Wrap(err, "eventPub.Publish")
			}

			return h.eventPub.Publish(ctx, bytes, "text/plain")
		},
	}

	_, err = h.store.CreateQuizTx(ctx, arg)
	if err != nil {
		return errors.Wrap(err, "store.CreateQuizTx")
	}

	return nil
}
