package llms

import "context"

type GenAI interface {
	GenerateContent(context.Context, string) (string, error)
}

type LLMApiKey string
