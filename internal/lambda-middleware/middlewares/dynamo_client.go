package middlewares

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/warpspeedboilerplate/go/internal/dynamo"
	lambdamiddleware "github.com/warpspeedboilerplate/go/internal/lambda-middleware"
)

var (
	ErrEnvironmentConfigNotFound = errors.New("environmentConfig error")
	ErrNotAuthorized             = errors.New("not authorized")
)

func DynamoClientMiddleware[Args any, Reply any](dynamoClient *dynamodb.Client) lambdamiddleware.Middleware[Args, Reply] {
	return func(ctx context.Context, invocation *lambdamiddleware.Chain[Args, Reply]) (*Reply, error) {
		if invocation.Event.Identity == nil {
			return nil, ErrNotAuthorized
		}

		if invocation.Event.Identity.CognitoIdentityAuthProvider == nil {
			return nil, ErrNotAuthorized
		}

		if dynamoClient == nil {
			client, err := dynamo.NewDynamoClientWithDefaultConfig(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to create dynamo client because %w", err)
			}

			invocation.DynamoClient = client
		} else {
			invocation.DynamoClient = dynamoClient
		}

		return nil, nil
	}
}
