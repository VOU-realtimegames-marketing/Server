package db

import (
	"context"
)

type CreateQuizTxParams struct {
	CreateQuizParams
	AfterCreate func(quiz Quiz) error
}

type CreateQuizTxResult struct {
	Quiz Quiz
}

func (store *SQLStore) CreateQuizTx(ctx context.Context, arg CreateQuizTxParams) (CreateQuizTxResult, error) {
	var result CreateQuizTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Quiz, err = q.CreateQuiz(ctx, arg.CreateQuizParams)
		if err != nil {
			return err
		}

		return arg.AfterCreate(result.Quiz)
	})

	return result, err
}
