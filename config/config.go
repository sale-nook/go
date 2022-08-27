package config

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/sethvargo/go-envconfig"

	"github.com/davemackintosh/aws-appsync-go/internal/types"
)

var (
	ErrUnknownEnvironment      = fmt.Errorf("unknown environment")
	ErrMissingUserPoolClientID = fmt.Errorf("missing UserPoolClientID. Is ./config/api.json empty?")
	ErrMissingUserPool         = fmt.Errorf("missing UserPool. Is ./config/api.json empty?")
)

type Config struct {
	// From outputs of cdk.
	UserPoolClientID *string
	UserPoolID       *string

	// From env variables.
	Environment *string `env:"ENVIRONMENT"`
	AWSRegion   *string `env:"AWS_REGION, required"`

	// Integrations.
	PlaidClientID     *string `env:"PLAID_CLIENT_ID, required"`
	PlaidSecret       *string `env:"PLAID_SECRET, required"`
	OAuthCallbackBase *string `env:"OAUTH_CALLBACK_ROOT,required"`

	// For cdk.
	GithubAccessToken *string `env:"GITHUB_ACCESS_TOKEN,required"`
}

//go:embed api.json
var content []byte

func GetConfig() (*Config, error) {
	var (
		config        Config
		cdkOutputEnvs types.CDKOutputsByEnv
		cdkOutputs    types.CDKOutputs
	)

	envErr := envconfig.Process(context.Background(), &config)
	if envErr != nil {
		return nil, fmt.Errorf("failed to process environments in env for config %w", envErr)
	}

	err := json.Unmarshal(content, &cdkOutputEnvs)
	if err != nil {
		return nil, fmt.Errorf("error during config Unmarshal: %w", err)
	}

	// Now we can get the config for the current Environment
	switch *config.Environment {
	case "ci":
		cdkOutputs = cdkOutputEnvs.Staging
	case "staging":
		cdkOutputs = cdkOutputEnvs.Staging
	case "production":
		cdkOutputs = cdkOutputEnvs.Production
	default:
		return nil, ErrUnknownEnvironment
	}

	if cdkOutputs.UserPoolClientID == nil {
		return nil, ErrMissingUserPoolClientID
	}

	if cdkOutputs.UserPoolID == nil {
		return nil, ErrMissingUserPool
	}

	config.UserPoolClientID = cdkOutputs.UserPoolClientID
	config.UserPoolID = cdkOutputs.UserPoolID

	return &config, nil
}
