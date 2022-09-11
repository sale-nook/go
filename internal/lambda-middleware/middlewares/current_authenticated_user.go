package middlewares

import (
	"context"
	"fmt"
	"strings"

	lambdamiddleware "github.com/davemackintosh/go/internal/lambda-middleware"
	"github.com/davemackintosh/go/internal/types"
	"github.com/davemackintosh/go/internal/utils"
)

var (
	ErrNoIdentityInEvent      = fmt.Errorf("no identity in event")
	ErrNoIdentityAuthProvider = fmt.Errorf("no identity auth provider")
	ErrMalformedAuthProvider  = fmt.Errorf("malformed auth provider")
)

// CurrentAuthenticatedUser is a middleware that extracts the current authenticated user from the event and sets it in the chain as. It does not query for said user but only extracts their possible user id.
func CurrentAuthenticatedUserID[Args any, Reply any](ctx context.Context, invocation *lambdamiddleware.Chain[Args, Reply]) (*Reply, error) {
	if invocation.Event.Identity == nil {
		return nil, ErrNoIdentityInEvent
	}

	if invocation.Event.Identity.CognitoIdentityAuthProvider == nil {
		return nil, ErrNoIdentityAuthProvider
	}

	authProviderParts := strings.Split(*invocation.Event.Identity.CognitoIdentityAuthProvider, ":")

	if len(authProviderParts) != 3 {
		return nil, ErrMalformedAuthProvider
	}

	if invocation.Auth == nil {
		invocation.Auth = &types.AppSyncLambdaIdentityEventIdentity{}
	}

	if authProviderParts[2] != "" {
		invocation.Auth.UserID = utils.ToPointer(strings.TrimSuffix(authProviderParts[2], "\""))
	}

	if invocation.Auth.UserID == nil {
		return nil, ErrMalformedAuthProvider
	}

	//nolint: nilnil
	return nil, nil
}
