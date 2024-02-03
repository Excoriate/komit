package ai

import "context"

type Config interface {
	GetBaseURL() string
	GetEndpointName() string
	GetMaxTokens() int
	GetPassword() string
	GetModel() string
	GetEngine() string
	GetTemperature() float32
	GetProviderRegion() string
	GetAuthToken() string
}

type Provider interface {
	Authenticate(authToken string) error
	Configure(config Config) error
	GetCompletion(ctx context.Context, prompt string) (string, error)
}

type AIProvider struct {
	Name           string
	Model          string
	Password       string
	BaseURL        string
	EndpointName   string
	Engine         string
	Temperature    float32
	ProviderRegion string
	MaxTokens      int
	AuthToken      string
}

func (p *AIProvider) GetBaseURL() string {
	return p.BaseURL
}

func (p *AIProvider) GetEndpointName() string {
	return p.EndpointName
}

func (p *AIProvider) GetMaxTokens() int {
	return p.MaxTokens
}

func (p *AIProvider) GetPassword() string {
	return p.Password
}

func (p *AIProvider) GetModel() string {
	return p.Model
}

func (p *AIProvider) GetEngine() string {
	return p.Engine
}
func (p *AIProvider) GetTemperature() float32 {
	return p.Temperature
}

func (p *AIProvider) GetProviderRegion() string {
	return p.ProviderRegion
}

func (p *AIProvider) GetAuthToken() string {
	return p.AuthToken
}
