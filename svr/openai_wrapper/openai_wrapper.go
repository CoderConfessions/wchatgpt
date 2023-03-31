package openaiwrapper

import (
	"context"
	"time"

	"github.com/sashabaranov/go-openai"
)

var token string

func SetupToken(token_ string) {
	token = token_
}

func newClient() *openai.Client {
	return openai.NewClient(token)
}

func SingleCompletion(prompt string) (openai.CompletionResponse, error) {
	ctx := context.TODO()
	subctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	return newClient().CreateCompletion(
		subctx,
		openai.CompletionRequest{
			Model:     openai.GPT3TextDavinci003,
			Prompt:    prompt,
			MaxTokens: 1024,
		},
	)
}
