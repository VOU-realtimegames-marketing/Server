package handler

import (
	"VOU-Server/pkg/rabbitmq/publisher"
	"context"

	"github.com/google/wire"
)

var _ QuizGenHandler = (*quizGenHandler)(nil)

var QuizGenHandlerSet = wire.NewSet(NewQuizGenHandler)

type quizGenHandler struct {
	eventPub publisher.EventPublisher
}

func NewQuizGenHandler(eventPub publisher.EventPublisher) QuizGenHandler {
	return &quizGenHandler{
		eventPub: eventPub,
	}
}

func (h *quizGenHandler) Handle(ctx context.Context, e any) error {
	// slog.Info("received event", "event.BaristaOrdered", e)

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
