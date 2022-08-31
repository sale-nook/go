package lambdamiddleware

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/davemackintosh/cdk-appsync-go/config"
	"github.com/davemackintosh/cdk-appsync-go/internal/types"
)

var ErrEnvironmentConfigNotFound = errors.New("environmentConfig error")

type Middleware[Args any, Reply any] func(context.Context, *Chain[Args, Reply]) (*Reply, error)

// A Chain is a collection of functions that return functions to be
// executed in the scope of a lambda invocation. The functions in the Chain
// are executed in the order they are added to the Chain.
type Chain[Args any, Response any] struct {
	Auth         *types.AppSyncLambdaIdentityEventIdentity
	EnvConfig    *config.Config
	Event        types.AppSyncLambdaIdentityEvent[Args]
	Middleware   []Middleware[Args, Response]
	Response     types.AppsyncLambdaResponseEvent[Response]
	DynamoClient *dynamodb.Client
}

// NewMiddlewareChain returns a new Chain.
func NewMiddlewareChain[Args any, Response any](ctx context.Context, event types.AppSyncLambdaIdentityEvent[Args]) *Chain[Args, Response] {
	return &Chain[Args, Response]{
		Event: event,
	}
}

// Use adds a middleware function to the Chain.
func (c *Chain[Args, Response]) Use(middleware ...Middleware[Args, Response]) *Chain[Args, Response] {
	c.Middleware = append(c.Middleware, middleware...)

	return c
}

// Invoke executes the Chain.
func (c *Chain[Args, Response]) Invoke(ctx context.Context) (*Response, error) {
	for _, middleware := range c.Middleware {
		response, err := middleware(ctx, c)
		if err != nil {
			return nil, err
		}

		c.Response.Response = response
	}

	return c.Response.Response, nil
}
