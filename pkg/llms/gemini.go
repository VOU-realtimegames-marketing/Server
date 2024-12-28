package llms

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"github.com/google/wire"
	"google.golang.org/api/option"
)

var GeminiSet = wire.NewSet(NewGemini)

func NewGenerativeAIClient(ctx context.Context, geminiApiKey LLMApiKey) (*genai.Client, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(string(geminiApiKey)))
	return client, err
}

type Gemini struct {
	model *genai.GenerativeModel
}

func NewGemini(genAIClient *genai.Client) GenAI {
	modelName := "gemini-1.5-flash"
	model := genAIClient.GenerativeModel(modelName)
	return &Gemini{
		model: model,
	}
}

func (g *Gemini) GenerateContent(ctx context.Context, prompt string) (string, error) {
	fmt.Println("Generated content:")
	resp, err := g.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	return fmt.Sprint(resp.Candidates[0].Content.Parts[0]), nil
}
