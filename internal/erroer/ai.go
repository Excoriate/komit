package erroer

import "fmt"

// ErrMaxTokensExceedApplicationLimit indicates that the max tokens exceed the application's limit.
type ErrMaxTokensExceedApplicationLimit struct {
	MaxTokens  int
	Model      string
	ErrWrapped error
}

func (e *ErrMaxTokensExceedApplicationLimit) Error() string {
	return fmt.Sprintf("max tokens %d exceed application limit: %s when using the model %s", e.MaxTokens, e.ErrWrapped.Error(), e.Model)
}

func (e *ErrMaxTokensExceedApplicationLimit) Unwrap() error {
	return e.ErrWrapped
}

// NewErrMaxTokensExceedApplicationLimit creates a new ErrMaxTokensExceedApplicationLimit error.
func NewErrMaxTokensExceedApplicationLimit(maxTokens int, model string, err error) *ErrMaxTokensExceedApplicationLimit {
	return &ErrMaxTokensExceedApplicationLimit{
		MaxTokens:  maxTokens,
		Model:      model,
		ErrWrapped: err,
	}
}

// ErrMaxTokensExceedModelLimit indicates that the max tokens exceed the model's limit.
type ErrMaxTokensExceedModelLimit struct {
	Model      string
	MaxTokens  int
	ErrWrapped error
}

func (e *ErrMaxTokensExceedModelLimit) Error() string {
	return fmt.Sprintf("max tokens %d exceed model limit: %s when using the model %s", e.MaxTokens, e.ErrWrapped.Error(), e.Model)
}

func (e *ErrMaxTokensExceedModelLimit) Unwrap() error {
	return e.ErrWrapped
}

// NewErrMaxTokensExceedModelLimit creates a new ErrMaxTokensExceedModelLimit error.
func NewErrMaxTokensExceedModelLimit(maxTokens int, model string, err error) *ErrMaxTokensExceedModelLimit {
	return &ErrMaxTokensExceedModelLimit{
		Model:      model,
		MaxTokens:  maxTokens,
		ErrWrapped: err,
	}
}

// ErrOpenAIAPIError indicates an error from the OpenAI API.
type ErrOpenAIAPIError struct {
	ErrWrapped error
	Details    string
}

func (e *ErrOpenAIAPIError) Error() string {
	return fmt.Sprintf("openai api error: %s, details: %s", e.ErrWrapped.Error(), e.Details)
}

func (e *ErrOpenAIAPIError) Unwrap() error {
	return e.ErrWrapped
}

// NewErrOpenAIAPIError creates a new ErrOpenAIAPIError.
func NewErrOpenAIAPIError(err error, details string) *ErrOpenAIAPIError {
	return &ErrOpenAIAPIError{
		ErrWrapped: err,
		Details:    details,
	}
}
