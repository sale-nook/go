package config

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/sethvargo/go-envconfig"

	"github.com/davemackintosh/cdk-appsync-go/internal/types"
)

var (
	ErrUnknownEnvironment = fmt.Errorf("unknown environment")
)

type Config struct {
	// From outputs of cdk.
	UserPoolClientID *string
	UserPoolID       *string

	// From env variables.
	Environment *string `env:"ENVIRONMENT"`
	AWSRegion   *string `env:"AWS_REGION, required"`

	// Integrations.
	OAuthCallbackBase *string `env:"OAUTH_CALLBACK_ROOT,required"`

	// For cdk.
	GithubAccessToken *string `env:"GITHUB_ACCESS_TOKEN,required"`
}

func GetConfig() (*Config, error) {
	var (
		config     Config
		cdkOutputs types.CDKOutputs
	)

	envErr := envconfig.Process(context.Background(), &config)
	if envErr != nil {
		return nil, fmt.Errorf("failed to process environments in env for config %w", envErr)
	}

	config.UserPoolClientID = cdkOutputs.UserPoolClientID
	config.UserPoolID = cdkOutputs.UserPoolID

	return &config, nil
}
