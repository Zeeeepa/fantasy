package bedrock

import (
	"fmt"

	"github.com/charmbracelet/fantasy/ai"
)

// https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters-anthropic-claude-text-completion.html
// https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters-anthropic-claude-messages.html

type claudeInput struct {
	AnthropicVersion string          `json:"anthropic_version"`
	MaxTokens        *int64          `json:"max_tokens"`
	Messages         []claudeMessage `json:"messages"`
}

type claudeMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func toClaudeInput(call ai.Call) (claudeInput, error) {
	var err error
	i := claudeInput{
		MaxTokens: call.MaxOutputTokens,
	}

	i.Messages, err = toClaudePrompt(call.Prompt)
	if err != nil {
		return i, err
	}
	return i, nil
}

func toClaudePrompt(prompt ai.Prompt) (messages []claudeMessage, err error) {
	for _, m := range prompt {
		message := claudeMessage{
			Role: string(m.Role),
		}

		for _, part := range m.Content {
			switch content := part.(type) {
			case ai.TextPart:
				message.Content = content.Text
			default:
				return nil, fmt.Errorf("fantasy: ")
			}
		}

		messages = append(messages, message)
	}
	return messages, err
}
