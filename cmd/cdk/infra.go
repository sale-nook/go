package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"

	"github.com/warpspeedboilerplate/go/cmd/cdk/internal/stacks"
	"github.com/warpspeedboilerplate/go/internal/utils"
)

func main() {
	app := awscdk.NewApp(nil)

	// Our application is broken up into different stacks.
	infra := stacks.NewInfraStack(app)
	apiStack, api := stacks.NewAppsyncStack(app, infra)
	stacks.NewDatabaseStack(app, infra)
	profileStack := stacks.NewProfileStack(app, api, infra)

	// Add dependencies.
	profileStack.
		AddDependency(apiStack, utils.ToPointer("appsync stack configures the user pool & client required for users."))

	app.Synth(nil)
}
