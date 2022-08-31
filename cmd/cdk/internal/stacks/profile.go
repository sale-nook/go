package stacks

import (
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"

	"github.com/davemackintosh/cdk-appsync-go/cmd/cdk/internal"
	"github.com/davemackintosh/cdk-appsync-go/cmd/cdk/internal/iam"
	"github.com/davemackintosh/cdk-appsync-go/cmd/cdk/internal/lambda"
	"github.com/davemackintosh/cdk-appsync-go/internal/cdkutils"
)

func NewProfileStack(app awscdk.App, api awsappsync.CfnGraphQLApi, infra *internal.InfraEntities) awscdk.Stack { //nolint:ireturn
	env, err := internal.InfraAccountAndRegion()
	if err != nil {
		log.Fatalf("failed to get infra account and region %s", err)
	}

	stackName := cdkutils.NameWithEnvironment("profile")
	stack := awscdk.NewStack(app, &stackName, &awscdk.StackProps{
		Env: env,
	})
	lambdaSources := []lambda.DataResolverProps{
		{
			CmdPath:   "get-profile",
			FieldName: "getProfile",
			TypeName:  lambda.DataResolverTypesQUERY,
			IAMPolicies: []awsiam.PolicyStatement{
				iam.GetPolicyStatementFor(stack, iam.PolicyNameDynamoUsersRead),
				iam.GetPolicyStatementFor(stack, iam.PolicyNameUseKMS),
			},
		},
	}

	// Create the lambdas, well assign roles next.
	lambda.CreateLambdasFromMap(stack, api, lambdaSources, infra)

	return stack
}
