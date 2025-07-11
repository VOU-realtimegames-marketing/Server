package llms

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"google.golang.org/genai"
)

var GeminiSet = wire.NewSet(NewGemini)

func NewGenerativeAIClient(ctx context.Context, geminiApiKey LLMApiKey) (*genai.Client, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  string(geminiApiKey),
		Backend: genai.BackendGeminiAPI,
	})
	return client, err
}

type Gemini struct {
	model     *genai.Models
	modelName string
}

func NewGemini(genAIClient *genai.Client) GenAI {
	modelName := "gemini-2.0-flash"
	model := genAIClient.Models
	return &Gemini{
		model:     model,
		modelName: modelName,
	}
}

func (g *Gemini) GenerateContent(ctx context.Context, prompt string, config any) (string, error) {

	cfg, ok := config.(*genai.GenerateContentConfig)
	if !ok {
		return "", fmt.Errorf("invalid config type: %T", config)
	}

	resp, err := g.model.GenerateContent(ctx, g.modelName, genai.Text(prompt), cfg)
	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
