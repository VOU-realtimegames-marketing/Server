package db

import (
	"context"
)

type CreateEventTxParams struct {
	CreateEventParams
	AfterCreate func(event Event) error
}

type CreateEventTxResult struct {
	Event Event
}

func (store *SQLStore) CreateEventTx(ctx context.Context, arg CreateEventTxParams) (CreateEventTxResult, error) {
	var result CreateEventTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Event, err = q.CreateEvent(ctx, arg.CreateEventParams)
		if err != nil {
			return err
		}

		return arg.AfterCreate(result.Event)
	})

	return result, err
}
