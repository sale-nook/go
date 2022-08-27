package cdkutils_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/davemackintosh/aws-appsync-go/internal"
	"github.com/davemackintosh/aws-appsync-go/internal/cdkutils"
)

func TestNameWithEnvironment(t *testing.T) {
	app := internal.GetApp()
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "ci"
	}

	assert.Equal(t, fmt.Sprintf("%s-%s-my-resource", app.Name, env), cdkutils.NameWithEnvironment("my-resource"))
	assert.Equal(t, fmt.Sprintf("%s-%s-my-stack-my-resource", app.Name, env), cdkutils.NameWithStackAndEnvironment(env+"-my-resource", "my-stack"))
	assert.Equal(t, fmt.Sprintf("%s-%s-my-stack-my-resource", app.Name, env), cdkutils.NameWithStackAndEnvironment(env+"-"+env+"-my-resource", "my-stack"))
}
