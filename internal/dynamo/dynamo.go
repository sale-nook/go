package dynamo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/davemackintosh/cdk-appsync-go/config"
)

func NewDynamoClientWithDefaultConfig(ctx context.Context) (*dynamodb.Client, error) {
	env, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	cfg, err := awsconfig.LoadDefaultConfig(ctx, func(o *awsconfig.LoadOptions) error {
		o.Region = *env.AWSRegion

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return NewDynamoClient(ctx, cfg), nil
}

func NewDynamoClient(ctx context.Context, config aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(config)
}
