package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/davemackintosh/aws-appsync-go/cmd/create-user-from-cognito/internal"
	"github.com/davemackintosh/aws-appsync-go/internal/dynamo/tables"
)

func NewUserFromCognitoEvent(event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, func(o *config.LoadOptions) error {
		o.Region = "eu-west-2"

		return nil
	})
	if err != nil {
		return event, fmt.Errorf("failed to load config: %w", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	err = internal.WriteUserToTableWithClient(ctx, client, tables.User{
		ID:        event.CognitoEventUserPoolsHeader.UserName,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return event, fmt.Errorf("failed to write user to table: %w", err)
	}

	return event, nil
}
