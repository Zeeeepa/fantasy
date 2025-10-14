package bedrock

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/charmbracelet/fantasy/ai"
)

type languageModel struct {
	provider string
	modelID  string
	client   *bedrockruntime.Client
}

func (b languageModel) Model() string {
	return b.modelID
}

func (b languageModel) Provider() string {
	return b.provider
}

func (b languageModel) Generate(ctx context.Context, call ai.Call) (*ai.Response, error) {
	params, err := b.prepareParams(call)
	if err != nil {
		return nil, err
	}

	output, err := b.client.InvokeModel(ctx, params)
	if err != nil {
		return nil, err
	}

	panic(fmt.Sprintf("bedrock output: %+v", output))

	// return &ai.Response{
	// 	Content: content,
	// 	// Usage: ai.Usage{
	// 	// 	InputTokens:         response.Usage.InputTokens,
	// 	// 	OutputTokens:        response.Usage.OutputTokens,
	// 	// 	TotalTokens:         response.Usage.InputTokens + response.Usage.OutputTokens,
	// 	// 	CacheCreationTokens: response.Usage.CacheCreationInputTokens,
	// 	// 	CacheReadTokens:     response.Usage.CacheReadInputTokens,
	// 	// },
	// 	// FinishReason:     mapFinishReason(string(response.StopReason)),
	// 	ProviderMetadata: ai.ProviderMetadata{},
	// 	// Warnings:         warnings,
	// }, nil
}

func (b languageModel) Stream(ctx context.Context, call ai.Call) (ai.StreamResponse, error) {
	return nil, errors.New("bedrock provider not fully implemented")
}

func (b languageModel) prepareParams(call ai.Call) (*bedrockruntime.InvokeModelInput, error) {
	input := bedrockruntime.InvokeModelInput{
		// ModelId:     ptr(fmt.Sprintf("us-east-1.%s", b.modelID)),
		ModelId: ptr(b.modelID),
		// ModelId:     ptr("us-east-1.anthropic.claude-sonnet-4-5-v2:0"),
		ContentType: ptr("application/json"),
		// Body:        body,
	}

	// call.Prompt

	switch {
	case containsAny(b.modelID, "anthropic", "claude", "sonnet"):
		i, err := toClaudeInput(call)
		if err != nil {
			return nil, err
		}
		input.Body, err = json.Marshal(i)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("fantasy: bedrock provider does not support model: %s", b.modelID)
	}

	return &input, nil
}
