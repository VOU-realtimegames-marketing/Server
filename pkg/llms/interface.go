package llms

import "context"

type GenAI interface {
	GenerateContent(context.Context, string, any) (string, error)
}

type LLMApiKey string
