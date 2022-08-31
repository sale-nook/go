package middlewares_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	lambdamiddleware "github.com/davemackintosh/cdk-appsync-go/internal/lambda-middleware"
	"github.com/davemackintosh/cdk-appsync-go/internal/lambda-middleware/middlewares"
	"github.com/davemackintosh/cdk-appsync-go/internal/types"
	"github.com/davemackintosh/cdk-appsync-go/internal/utils"
)

type TestResponse struct {
	Message string
}

func Test_CurrentAuthenticatedUserID(t *testing.T) {
	ctx := context.TODO()
	event := types.AppSyncLambdaIdentityEvent[any]{
		Identity: &types.AppSyncLambdaIdentityEventIdentity{
			CognitoIdentityAuthProvider: utils.ToPointer("test:123:userID1234"),
		},
	}
	middleware := lambdamiddleware.NewMiddlewareChain[any, TestResponse](ctx, event)
	assert.NotNil(t, middleware)

	reply, err := middleware.
		Use(middlewares.CurrentAuthenticatedUserID[any, TestResponse]).
		Use(func(ctx context.Context, invocation *lambdamiddleware.Chain[any, TestResponse]) (*TestResponse, error) {
			assert.NotNil(t, invocation.Auth.UserID)

			return &TestResponse{
				Message: "Hello",
			}, nil
		}).
		Invoke(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, reply)
	assert.Equal(t, "userID1234", *middleware.Auth.UserID)
}

func Test_CurrentAuthenticatedUserIDMissingIdentity(t *testing.T) {
	ctx := context.TODO()
	event := types.AppSyncLambdaIdentityEvent[any]{}
	middleware := lambdamiddleware.NewMiddlewareChain[any, TestResponse](ctx, event)
	assert.NotNil(t, middleware)

	reply, err := middleware.
		Use(middlewares.CurrentAuthenticatedUserID[any, TestResponse]).
		Invoke(ctx)

	assert.Error(t, err)
	assert.Nil(t, reply)
}

func Test_CurrentAuthenticatedUserIDMalformedProvider(t *testing.T) {
	ctx := context.TODO()
	middleware := lambdamiddleware.NewMiddlewareChain[any, TestResponse](ctx, types.AppSyncLambdaIdentityEvent[any]{
		Identity: &types.AppSyncLambdaIdentityEventIdentity{
			CognitoIdentityAuthProvider: utils.ToPointer("test:123"),
		},
	})
	assert.NotNil(t, middleware)

	reply, err := middleware.
		Use(middlewares.CurrentAuthenticatedUserID[any, TestResponse]).
		Invoke(ctx)

	assert.Error(t, err)
	assert.Nil(t, reply)

	middleware = lambdamiddleware.NewMiddlewareChain[any, TestResponse](ctx, types.AppSyncLambdaIdentityEvent[any]{
		Identity: &types.AppSyncLambdaIdentityEventIdentity{
			CognitoIdentityAuthProvider: utils.ToPointer("test:123:"),
		},
	})
	assert.NotNil(t, middleware)

	reply, err = middleware.
		Use(middlewares.CurrentAuthenticatedUserID[any, TestResponse]).
		Invoke(ctx)

	assert.Error(t, err)
	assert.Nil(t, reply)
}
