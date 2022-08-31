package internal_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/davemackintosh/cdk-appsync-go/cmd/cdk/internal"
)

func Test_InfraAccountAndRegion(t *testing.T) {
	os.Setenv("AWS_ACCOUNT_ID", "account") // nolint: tenv
	os.Setenv("AWS_REGION", "region")      // nolint: tenv

	env, err := internal.InfraAccountAndRegion()
	assert.NoError(t, err)
	assert.Equal(t, "account", *env.Account)
	assert.Equal(t, "region", *env.Region)

	os.Unsetenv("AWS_ACCOUNT_ID")

	assert.Panics(t, func() { internal.InfraAccountAndRegion() })
	os.Unsetenv("AWS_REGION")
	assert.Panics(t, func() { internal.InfraAccountAndRegion() })
}
