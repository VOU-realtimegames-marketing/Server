package handler

import (
	"VOU-Server/internal/pkg/task"
	"context"
)

type QuizGenHandler interface {
	Handle(context.Context, task.PayloadGenQuiz) error
}
