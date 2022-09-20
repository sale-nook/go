package config

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

var (
	ErrUnknownEnvironment = fmt.Errorf("unknown environment")
)

type Config struct {
	Environment       *string `env:"ENVIRONMENT"`
	AWSRegion         *string `env:"AWS_REGION, required"`
	OAuthCallbackBase *string `env:"OAUTH_CALLBACK_ROOT"`
	GithubAccessToken *string `env:"GITHUB_ACCESS_TOKEN"`
}

func GetConfig() (*Config, error) {
	var config Config

	envErr := envconfig.Process(context.Background(), &config)
	if envErr != nil {
		return nil, fmt.Errorf("failed to process environments in env for config %w", envErr)
	}

	return &config, nil
}
