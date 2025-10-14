package providertests

import (
	"net/http"
	"testing"

	"github.com/charmbracelet/fantasy/ai"
	"github.com/charmbracelet/fantasy/bedrock"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

// const defaultBaseURL = "https://fantasy-playground-resource.services.ai.azure.com/"

func TestBedrockCommon(t *testing.T) {
	testCommon(t, []builderPair{
		{"bedrock-anthropic-claude-v2", builderBedrockClaudeV2(t), nil},
	})
}

// func TestBedrockThinking(t *testing.T) {
// 	opts := ai.ProviderOptions{
// 		bedrock.Name: &bedrock.ProviderOptions{
// 			ReasoningEffort: openai.ReasoningEffortOption(openai.ReasoningEffortLow),
// 		},
// 	}
// 	testThinking(t, []builderPair{
// 		{"bedrock-anthropic-claude-v2", builderBedrockClaudeV2(t), opts},
// 	}, testBedrockThinking)
// }

// func testBedrockThinking(t *testing.T, result *ai.AgentResult) {
// 	require.Greater(t, result.Response.Usage.ReasoningTokens, int64(0), "expected reasoning tokens, got none")
// }

func builderBedrockClaudeV2(t *testing.T) func(r *recorder.Recorder) (ai.LanguageModel, error) {
	return func(r *recorder.Recorder) (ai.LanguageModel, error) {
		provider, err := bedrock.New(
			t.Context(),
			bedrock.WithHTTPClient(&http.Client{Transport: r}),
		)
		if err != nil {
			return nil, err
		}
		// return provider.LanguageModel("anthropic.claude-sonnet-4-5-20250929-v1:0")
		return provider.LanguageModel("us.anthropic.claude-3-sonnet-20240229-v1:0")
	}
}
