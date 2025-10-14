package bedrock

import "github.com/charmbracelet/fantasy/ai"

type ProviderOptions struct {
	// Add Bedrock-specific options here
}

func (o *ProviderOptions) Options() {}

func NewProviderOptions(opts *ProviderOptions) ai.ProviderOptions {
	return ai.ProviderOptions{
		Name: opts,
	}
}
