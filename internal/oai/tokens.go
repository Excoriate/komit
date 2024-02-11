package oai

import (
	"fmt"

	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"

	"log"
	"strings"
)

// MaxTokensMap maps the OpenAI model names to their maximum token limits.
var MaxTokensMap = map[string]int{
	"gpt-4-turbo-preview":    128000, // Assuming based on the context; adjust as necessary.
	"gpt-4-0125-preview":     128000, // Assuming based on the context; adjust as necessary.
	"gpt-4-1106-preview":     128000, // Assuming based on the context; adjust as necessary.
	"gpt-4-vision-preview":   128000, // Assuming based on the context; adjust as necessary.
	"gpt-4":                  8192,
	"gpt-3.5-turbo-0125":     16385,
	"gpt-3.5-turbo":          4096,
	"gpt-3.5-turbo-instruct": 4096,
	"gpt-3.5-turbo-1106":     16385,
}

// GetMaxTokens returns the maximum token limit for the given model.
// If the model is not recognized, it returns -1.
func GetMaxTokens(model string) int {
	if maxTokens, exists := MaxTokensMap[model]; exists {
		return maxTokens
	}
	return -1
}

// NumTokensFromMessages OpenAI Cookbook: https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb
func NumTokensFromMessages(messages []openai.ChatCompletionMessage, model string) (numTokens int) {
	tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		err = fmt.Errorf("encoding for model: %v", err)
		log.Println(err)
		return
	}

	var tokensPerMessage, tokensPerName int
	switch model {
	case "gpt-3.5-turbo-0613",
		"gpt-3.5-turbo-16k-0613",
		"gpt-4-0314",
		"gpt-4-32k-0314",
		"gpt-4-0613",
		"gpt-4-32k-0613":
		tokensPerMessage = 3
		tokensPerName = 1
	case "gpt-3.5-turbo-0301":
		tokensPerMessage = 4 // every message follows <|start|>{role/name}\n{content}<|end|>\n
		tokensPerName = -1   // if there's a name, the role is omitted
	default:
		if strings.Contains(model, "gpt-3.5-turbo") {
			log.Println("warning: gpt-3.5-turbo may update over time. Returning num tokens assuming gpt-3.5-turbo-0613.")
			return NumTokensFromMessages(messages, "gpt-3.5-turbo-0613")
		} else if strings.Contains(model, "gpt-4") {
			log.Println("warning: gpt-4 may update over time. Returning num tokens assuming gpt-4-0613.")
			return NumTokensFromMessages(messages, "gpt-4-0613")
		} else {
			err = fmt.Errorf("num_tokens_from_messages() is not implemented for model %s. See https://github.com/openai/openai-python/blob/main/chatml.md for information on how messages are converted to tokens.", model)
			log.Println(err)
			return
		}
	}

	for _, message := range messages {
		numTokens += tokensPerMessage
		numTokens += len(tkm.Encode(message.Content, nil, nil))
		numTokens += len(tkm.Encode(message.Role, nil, nil))
		numTokens += len(tkm.Encode(message.Name, nil, nil))
		if message.Name != "" {
			numTokens += tokensPerName
		}
	}
	numTokens += 3 // every reply is primed with <|start|>assistant<|message|>
	return numTokens
}
