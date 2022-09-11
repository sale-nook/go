package middlewares

import (
	"context"
	"fmt"

	"github.com/davemackintosh/go/config"
	lambdamiddleware "github.com/davemackintosh/go/internal/lambda-middleware"
)

func EnvironmentConfig[Args any, Reply any](ctx context.Context, invocation *lambdamiddleware.Chain[Args, Reply]) (*Reply, error) {
	env, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	invocation.EnvConfig = env

	//nolint: nilnil
	return nil, nil
}
