package oai

import (
	"context"

	"github.com/excoriate/komit/internal/ai"
	"github.com/sashabaranov/go-openai"
)

type Client struct {
	client *openai.Client
}

func (c *Client) Authenticate(authToken string) error {
	return nil
}

func (c *Client) Configure(config ai.Config) error {
	return nil
}

func (c *Client) GetCompletion(ctx context.Context, prompt string) (string, error) {
	return "", nil
}

func NewOpenAI(config ai.Config) ai.Provider {
	cfg := openai.DefaultConfig(config.GetAuthToken())
	return &Client{
		client: openai.NewClientWithConfig(cfg),
	}
}
