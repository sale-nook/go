package internal_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/warpspeedboilerplate/go/cmd/cdk/internal"
)

func Test_InfraAccountAndRegion(t *testing.T) {
	os.Setenv("AWS_ACCOUNT_ID", "account") // nolint: tenv
	os.Setenv("AWS_REGION", "region")      // nolint: tenv

	env, err := internal.InfraAccountAndRegion()
	assert.NoError(t, err)
	assert.Equal(t, "account", *env.Account)
	assert.Equal(t, "region", *env.Region)

	os.Unsetenv("AWS_ACCOUNT_ID")

	_, err = internal.InfraAccountAndRegion()

	assert.NotNil(t, err)
	os.Unsetenv("AWS_REGION")

	_, err = internal.InfraAccountAndRegion()

	assert.NotNil(t, err)
}
