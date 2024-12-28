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

// sudo rabbitmqctl list_queues name messages_ready messages_unacknowledged
type Quiz struct {
	Question string   `json:"question"`
	Answer   string   `json:"answer"`
	Options  []string `json:"options"`
}

func (h *quizGenHandler) Handle(ctx context.Context, payload task.PayloadGenQuiz) error {
	log.Info().Msg("received event: quiz event generated")

	fmt.Println("payload: ", payload)

	quizzesStr, err := h.gemini.GenerateContent(ctx, fmt.Sprintf(task.QuizGeneratePrompt, 3, payload.QuizGenre))
	jsonStr := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(quizzesStr, "```json"), "```\n"))
	fmt.Println("jsonStr: ", jsonStr)

	var quizzes []Quiz
	err = json.Unmarshal([]byte(jsonStr), &quizzes)
	fmt.Println("quizzes: ", quizzes)
	h.store.CreateQuiz(ctx, db.CreateQuizParams{
		EventID:   payload.EventId,
		QuizGenre: payload.QuizGenre,
		Content:   []byte(jsonStr),
	})

	rspPayload := task.PayloadQuizCreated{
		Status: true,
	}
	bytes, err := json.Marshal(rspPayload)
	if err != nil {
		return errors.Wrap(err, "eventPub.Publish")
	}

	h.eventPub.Publish(ctx, bytes, "text/plain")

	// order := domain.NewBaristaOrder(e)

	// db := h.pg.GetDB()
	// querier := postgresql.New(db)

	// tx, err := db.Begin()
	// if err != nil {
	// 	return errors.Wrap(err, "baristaOrderedEventHandler.Handle")
	// }

	// qtx := querier.WithTx(tx)

	// _, err = qtx.CreateOrder(ctx, postgresql.CreateOrderParams{
	// 	ID:       order.ID,
	// 	ItemType: int32(order.ItemType),
	// 	ItemName: order.ItemName,
	// 	TimeUp:   order.TimeUp,
	// 	Created:  order.Created,
	// 	Updated: sql.NullTime{
	// 		Time:  order.Updated,
	// 		Valid: true,
	// 	},
	// })
	// if err != nil {
	// 	slog.Info("failed to call to repo", "error", err)

	// 	return errors.Wrap(err, "baristaOrderedEventHandler-querier.CreateOrder")
	// }

	// // todo: it might cause dual-write problem, but we accept it temporary
	// for _, event := range order.DomainEvents() {
	// 	eventBytes, err := json.Marshal(event)
	// 	if err != nil {
	// 		return errors.Wrap(err, "json.Marshal[event]")
	// 	}

	// 	if err := h.counterPub.Publish(ctx, eventBytes, "text/plain"); err != nil {
	// 		return errors.Wrap(err, "counterPub.Publish")
	// 	}
	// }

	// return tx.Commit()
	return nil
}
