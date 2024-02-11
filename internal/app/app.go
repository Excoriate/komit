package app

import (
	"context"
	"fmt"
	"strconv"

	"github.com/excoriate/komit/internal/oai"
	"github.com/excoriate/komit/pkg/env"

	"github.com/excoriate/komit/internal/ai"

	"github.com/excoriate/komit/pkg/logger"
)

type App struct {
	Ctx        context.Context
	Generate   Generate
	Log        logger.Log
	AIProvider ai.Provider
}

var (
	// TODO Implement in the near future?
	// providerOLlama = "ollama"
	defaultAIProvider = "openai"
)

type AIProviderOptions struct {
	Name      string
	AuthToken string
	Model     string
}

// New creates a new instance of the App
func New(ctx context.Context, aiProvider *AIProviderOptions) (*App, error) {
	a := &App{
		Ctx: ctx,
		Log: logger.NewLogger().Logger,
	}

	var provider ai.Provider

	// This is the limit that I'm imposing through the CLI ;)
	maxTokensApp, _ := strconv.Atoi(env.GetOrDefault("KOMIT_MAX_TOKENS", "2048"))

	// Setting AI provider configuration.
	aiProviderCfg := &ai.AIProvider{
		Name:      aiProvider.Name,
		AuthToken: aiProvider.AuthToken,
		Model:     aiProvider.Model,
		MaxTokens: maxTokensApp,
	}

	// Configuring the AI provider.
	if aiProvider.Name == defaultAIProvider {
		provider = oai.NewOpenAI(aiProviderCfg)
	} else {
		a.Log.Debug("provider %s not supported", aiProvider.Name)
		return nil, fmt.Errorf("provider %s not supported", aiProvider.Name)
	}

	a.AIProvider = provider
	a.Generate = NewGenerate(a)

	return a, nil
}
