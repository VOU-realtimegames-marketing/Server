package handler

import (
	"VOU-Server/internal/pkg/task"
	"context"
)

type QuizCreatedHandler interface {
	Handle(context.Context, task.PayloadQuizCreated) error
}
