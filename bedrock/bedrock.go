package bedrock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/charmbracelet/fantasy/ai"
)

const (
	Name = "bedrock"
	// DefaultURL = "https://bedrock-runtime.amazonaws.com"
)

type options struct {
	name       string
	httpClient bedrockruntime.HTTPClient

	// region       string
	// accessKey    string
	// secretKey    string
	// sessionToken string
	// endpoint     string
}

type provider struct {
	options options
	client  *bedrockruntime.Client
}

type Option = func(*options)

func New(ctx context.Context, opts ...Option) (ai.Provider, error) {
	providerOptions := options{
		name: Name,
	}
	for _, o := range opts {
		o(&providerOptions)
	}

	cfg, err := config.LoadDefaultConfig(ctx) //, config.WithRegion(providerOptions.region))
	if err != nil {
		return nil, fmt.Errorf("fantasy: unable to load default aws config: %w", err)
	}

	// if providerOptions.accessKey != "" && providerOptions.secretKey != "" {
	// 	cfg.Credentials = aws.CredentialsProviderFunc(
	// 		func(ctx context.Context) (aws.Credentials, error) {
	// 			return aws.Credentials{
	// 				AccessKeyID:     providerOptions.accessKey,
	// 				SecretAccessKey: providerOptions.secretKey,
	// 				SessionToken:    providerOptions.sessionToken,
	// 			}, nil
	// 		},
	// 	)
	// }

	client := bedrockruntime.NewFromConfig(
		cfg,
		func(o *bedrockruntime.Options) {
			if providerOptions.httpClient != nil {
				o.HTTPClient = providerOptions.httpClient
			}
		},
	)

	return &provider{
		options: providerOptions,
		client:  client,
	}, nil
}

func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

func WithHTTPClient(httpClient bedrockruntime.HTTPClient) Option {
	return func(o *options) {
		o.httpClient = httpClient
	}
}

// func WithRegion(region string) Option {
// 	return func(o *options) {
// 		o.region = region
// 	}
// }

// func WithCredentials(accessKey, secretKey, sessionToken string) Option {
// 	return func(o *options) {
// 		o.accessKey = accessKey
// 		o.secretKey = secretKey
// 		o.sessionToken = sessionToken
// 	}
// }

// func WithEndpoint(endpoint string) Option {
// 	return func(o *options) {
// 		o.endpoint = endpoint
// 	}
// }

func (b *provider) Name() string {
	return Name
}

func (b *provider) LanguageModel(modelID string) (ai.LanguageModel, error) {
	return languageModel{
		modelID:  modelID,
		provider: b.options.name,
		client:   b.client,
	}, nil
}

func (b *provider) ParseOptions(data map[string]any) (ai.ProviderOptionsData, error) {
	var options ProviderOptions
	if err := ai.ParseOptions(data, &options); err != nil {
		return nil, err
	}
	return &options, nil
}
