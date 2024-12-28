package handler

import "context"

type QuizCreatedHandler interface {
	Handle(context.Context, any) error
}
