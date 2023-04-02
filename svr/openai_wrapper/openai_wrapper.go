package openaiwrapper

import (
	"context"
	"time"

	"github.com/sashabaranov/go-openai"
)

var config openai.ClientConfig

func SetupOpenAIClientConfig(token string, baseURL string) {
	config = openai.DefaultConfig(token)
	if len(baseURL) != 0 {
		config.BaseURL = baseURL
	}
}

func newClient() *openai.Client {
	return openai.NewClientWithConfig(config)
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

func ChatCompletion(historyMessages []openai.ChatCompletionMessage, prompt string) (openai.ChatCompletionResponse, error) {
	ctx := context.TODO()
	subctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	messages := append(historyMessages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: prompt,
	})
	return newClient().CreateChatCompletion(
		subctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)
}
