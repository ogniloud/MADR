package study

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
)

const prompt = "Generate a paragraph using the following words to show their usage in context: "

type Generator struct {
	model *genai.GenerativeModel
}

func New(model *genai.GenerativeModel) Generator {
	return Generator{
		model: model,
	}
}

func (t Generator) GenerateText(ctx context.Context, words []string) (string, error) {
	req := prompt + strings.Join(words, ", ")

	resp, err := t.model.GenerateContent(ctx, genai.Text(req))
	if err != nil {
		return "", fmt.Errorf("Generator.GenerateText: %w", err)
	}

	return resp.Candidates[0].Content.Role, nil
}
