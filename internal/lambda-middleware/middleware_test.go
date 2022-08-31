package lambdamiddleware_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	lambdamiddleware "github.com/davemackintosh/cdk-appsync-go/internal/lambda-middleware"
	"github.com/davemackintosh/cdk-appsync-go/internal/types"
	"github.com/davemackintosh/cdk-appsync-go/internal/utils"
)

type TestReply struct {
	Message string
}

var (
	ErrEnvironmentConfigNotFound = errors.New("EnvironmentConfig error")
	ErrCurrentAuthenticatedUser  = errors.New("CurrentAuthenticatedUser error")
)

func TestGoodNewMiddlewareChain(t *testing.T) {
	ctx := context.TODO()
	event := types.AppSyncLambdaIdentityEvent[any]{}
	middleware := lambdamiddleware.NewMiddlewareChain[any, TestReply](ctx, event)
	assert.NotNil(t, middleware)

	reply, err := middleware.
		Use(func(ctx context.Context, invocation *lambdamiddleware.Chain[any, TestReply]) (*TestReply, error) {
			return &TestReply{
				Message: "Hello",
			}, nil
		}).
		Invoke(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, reply)

	assert.Equal(t, "Hello", reply.Message)
}

func TestBadEnvironmentNewMiddlewareChain(t *testing.T) {
	var EnvironmentConfig lambdamiddleware.Middleware[any, TestReply] = func(ctx context.Context, invocation *lambdamiddleware.Chain[any, TestReply]) (*TestReply, error) {
		return nil, ErrEnvironmentConfigNotFound
	}

	ctx := context.TODO()
	event := types.AppSyncLambdaIdentityEvent[any]{}
	middleware := lambdamiddleware.NewMiddlewareChain[any, TestReply](ctx, event)
	assert.NotNil(t, middleware)

	reply, err := middleware.
		Use(EnvironmentConfig).
		Invoke(ctx)

	assert.Error(t, err)
	assert.Nil(t, reply)

	assert.Equal(t, "EnvironmentConfig error", err.Error())
}

func TestMissingAuthNewMiddlewareChain(t *testing.T) {
	var CurrentAuthenticatedUser lambdamiddleware.Middleware[any, TestReply] = func(ctx context.Context, invocation *lambdamiddleware.Chain[any, TestReply]) (*TestReply, error) {
		return nil, ErrCurrentAuthenticatedUser
	}

	ctx := context.TODO()
	event := types.AppSyncLambdaIdentityEvent[any]{
		Identity: &types.AppSyncLambdaIdentityEventIdentity{
			CognitoIdentityID: utils.ToPointer("test:123"),
		},
	}
	middleware := lambdamiddleware.NewMiddlewareChain[any, TestReply](ctx, event)
	assert.NotNil(t, middleware)

	reply, err := middleware.
		Use(CurrentAuthenticatedUser).
		Invoke(ctx)

	assert.Error(t, err)
	assert.Nil(t, reply)

	assert.Equal(t, "CurrentAuthenticatedUser error", err.Error())
}
