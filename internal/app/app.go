package app

import (
	"context"
	"fmt"

	"github.com/excoriate/komit/internal/oai"

	"github.com/excoriate/komit/internal/ai"

	"github.com/excoriate/komit/pkg/logger"
)

type App struct {
	ctx      context.Context
	l        logger.Log
	provider ai.Provider
}

var (
	// TODO Implement in the near future?
	// providerOLlama = "ollama"
	providerOpenAI = "openai"
)

type AIProviderOptions struct {
	Name  string
	Token string
}

// New creates a new instance of the App
func New(ctx context.Context, aiProvider *AIProviderOptions) (*App, error) {
	logAdapter := logger.NewLogger()
	var provider ai.Provider

	// Setting AI provider configuration.
	aiProviderCfg := &ai.AIProvider{
		Name:      aiProvider.Name,
		AuthToken: aiProvider.Token,
	}

	// Configuring the AI provider.
	if aiProvider.Name == providerOpenAI {
		provider = oai.NewOpenAI(aiProviderCfg)
	} else {
		logAdapter.Logger.Debug("provider %s not supported", aiProvider.Name)
		return nil, fmt.Errorf("provider %s not supported", aiProvider.Name)
	}

	return &App{
		ctx:      ctx,
		l:        logAdapter.Logger,
		provider: provider,
	}, nil
}
