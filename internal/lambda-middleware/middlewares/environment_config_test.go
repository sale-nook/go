package middlewares_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	lambdamiddleware "github.com/davemackintosh/go/internal/lambda-middleware"
	"github.com/davemackintosh/go/internal/lambda-middleware/middlewares"
	"github.com/davemackintosh/go/internal/types"
)

type TestReply struct {
	Message string
}

func Test_EnvironmentConfig(t *testing.T) {
	ctx := context.TODO()
	event := types.AppSyncLambdaIdentityEvent[any]{}
	middleware := lambdamiddleware.NewMiddlewareChain[any, TestReply](ctx, event)
	assert.NotNil(t, middleware)

	reply, err := middleware.
		Use(middlewares.EnvironmentConfig[any, TestReply]).
		Use(func(ctx context.Context, invocation *lambdamiddleware.Chain[any, TestReply]) (*TestReply, error) {
			assert.NotNil(t, invocation.EnvConfig)

			return &TestReply{
				Message: "Hello",
			}, nil
		}).
		Invoke(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, reply)
	assert.Equal(t, "Hello", reply.Message)
}
