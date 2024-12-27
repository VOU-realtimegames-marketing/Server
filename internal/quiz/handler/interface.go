package handler

import "context"

type QuizGenHandler interface {
	Handle(context.Context, any) error
}
