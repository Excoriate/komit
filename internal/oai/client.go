package oai

import (
	"context"
	"errors"
	"fmt"

	"github.com/excoriate/komit/internal/erroer"

	"github.com/excoriate/komit/internal/ai"
	"github.com/sashabaranov/go-openai"
)

type Client struct {
	client           *openai.Client
	model            string
	temperature      float32
	maxToken         int
	presencePenalty  float32
	frequencyPenalty float32
}

const (
	// OpenAI completion parameters
	presencePenalty  = 0.0
	frequencyPenalty = 0.0
	topP             = 1.0
)

func (c *Client) Authenticate(authToken string) error {
	return nil
}

func (c *Client) Configure(config ai.Config) error {
	defaultConfig := openai.DefaultConfig(config.GetAuthToken())

	baseURL := config.GetBaseURL()
	if baseURL != "" {
		defaultConfig.BaseURL = baseURL
	}

	client := openai.NewClientWithConfig(defaultConfig)

	if client == nil {
		return errors.New("error creating OpenAI client")
	}

	c.client = client
	c.model = config.GetModel()
	c.temperature = config.GetTemperature()

	return nil
}

func (c *Client) GetCompletion(ctx context.Context, prompt string) (string, error) {
	if prompt == "" {
		return "", errors.New("prompt is empty")
	}

	msg := []openai.ChatCompletionMessage{
		{
			Role:    "user",
			Content: prompt,
		},
	}

	// Get an estimation of the tokens that I'll generate before I send the request to Open AI.
	estimatedTokens := NumTokensFromMessages(msg, c.model)
	// Get the maximum tokens per model
	maxTokenPerModel := GetMaxTokens(c.model)
	// The maximum token that we're imposing as part of the configuration
	maxTokenAllowed := c.maxToken

	if estimatedTokens > maxTokenAllowed {
		return "", erroer.
			NewErrMaxTokensExceedApplicationLimit(maxTokenAllowed, c.model,
				fmt.Errorf("prompt is too long, failed to pass the 'max' limit set at the application level"))
	}

	if estimatedTokens > maxTokenPerModel {
		return "", erroer.
			NewErrMaxTokensExceedModelLimit(maxTokenPerModel, c.model,
				fmt.Errorf("prompt is too long, failed to pass the 'max' limit that's supported by the model used"))
	}

	// Create a completion request
	resp, err := c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: c.model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature:      c.temperature,
		MaxTokens:        c.maxToken,
		PresencePenalty:  c.presencePenalty,
		FrequencyPenalty: c.frequencyPenalty,
		TopP:             topP,
	})
	if err != nil {
		return "", erroer.
			NewErrOpenAIAPIError(err, fmt.Sprintf("error returned from API while using the Chat Completition api with MaxTokens: %d, Model: %s, PresencePenalty: %f, FrequencyPenalty: %f, TopP: %f", c.maxToken, c.model, c.presencePenalty, c.frequencyPenalty, topP))
	}

	return resp.Choices[0].Message.Content, nil
}

func NewOpenAI(config ai.Config) ai.Provider {
	cfg := openai.DefaultConfig(config.GetAuthToken())

	return &Client{
		client:           openai.NewClientWithConfig(cfg),
		maxToken:         config.GetMaxTokens(),
		presencePenalty:  presencePenalty,
		model:            config.GetModel(),
		frequencyPenalty: frequencyPenalty,
	}
}
