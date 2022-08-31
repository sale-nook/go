package internal

import (
	"context"
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/sethvargo/go-envconfig"
)

type EnvironmentConfig struct {
	AccountID           *string    `json:"accountId" env:"AWS_ACCOUNT_ID, required"`
	Region              *string    `json:"region" env:"AWS_REGION, required"`
	OAuthCallbackRoot   *string    `json:"oauthCallbackRoot" env:"OAUTH_CALLBACK_ROOT, required"`
	CognitoCallbackURLS *[]*string `json:"cognitoCallbackUrls" env:"COGNITO_CALLBACK_URLS, required"`
}

func GetConfig() (*EnvironmentConfig, error) {
	var config EnvironmentConfig

	envErr := envconfig.Process(context.Background(), &config)
	if envErr != nil {
		return nil, fmt.Errorf("failed to process environments in env for cdk config %w", envErr)
	}

	return &config, nil
}

func InfraAccountAndRegion() (*awscdk.Environment, error) {
	config, err := GetConfig()
	if err != nil {
		return nil, err
	}

	return &awscdk.Environment{
		Account: config.AccountID,
		Region:  config.Region,
	}, nil
}
