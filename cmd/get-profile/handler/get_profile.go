package handler

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/davemackintosh/aws-appsync-go/internal/dal"
	"github.com/davemackintosh/aws-appsync-go/internal/dynamo"
	"github.com/davemackintosh/aws-appsync-go/internal/dynamo/tables"
	lambdamiddleware "github.com/davemackintosh/aws-appsync-go/internal/lambda-middleware"
	"github.com/davemackintosh/aws-appsync-go/internal/lambda-middleware/middlewares"
	wikitypes "github.com/davemackintosh/aws-appsync-go/internal/types"
	"github.com/davemackintosh/aws-appsync-go/internal/utils"
)

func GetProfile(ctx context.Context, event wikitypes.AppSyncLambdaIdentityEvent[any]) (*tables.User, error) {
	//nolint: wrapcheck
	return lambdamiddleware.NewMiddlewareChain[any, tables.User](ctx, event).
		Use(middlewares.CurrentAuthenticatedUserID[any, tables.User]).
		Use(func(ctx context.Context, invocation *lambdamiddleware.Chain[any, tables.User]) (*tables.User, error) {
			client, err := dynamo.NewDynamoClientWithDefaultConfig(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to create dynamo client because %w", err)
			}

			fmt.Printf("Searching for user with id %s", *invocation.Auth.UserID)
			user, getUserErr := client.GetItem(ctx, &dynamodb.GetItemInput{
				TableName: utils.ToPointer(dal.Tables().AWSAppSyncGoUser()),
				Key: map[string]types.AttributeValue{
					"id": &types.AttributeValueMemberS{Value: *invocation.Auth.UserID},
				},
			})

			if getUserErr != nil {
				return nil, fmt.Errorf("failed to get user: %w", getUserErr)
			}

			var userObj tables.User
			err = attributevalue.UnmarshalMap(user.Item, &userObj)

			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal user because %w", err)
			}

			return &userObj, nil
		}).
		Invoke(ctx)
}
