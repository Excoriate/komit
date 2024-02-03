package oai

import (
	"net/http"
)

const (
	openaiAPIURLv1                 = "https://api.openai.com/v1"
	defaultEmptyMessagesLimit uint = 300
)

type APIType string

const (
	APITypeOpenAI APIType = "OPEN_AI"
)

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	authToken string

	BaseURL              string
	OrgID                string
	APIType              APIType
	APIVersion           string                    // required when APIType is APITypeAzure or APITypeAzureAD
	AzureModelMapperFunc func(model string) string // replace model to azure deployment name func
	HTTPClient           *http.Client

	EmptyMessagesLimit uint
}

func DefaultConfig(authToken string) ClientConfig {
	return ClientConfig{
		authToken: authToken,
		BaseURL:   openaiAPIURLv1,
		APIType:   APITypeOpenAI,
		OrgID:     "",

		HTTPClient: &http.Client{},

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}

func (ClientConfig) String() string {
	return "<OpenAI API ClientConfig>"
}
