package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"

	"github.com/davemackintosh/aws-appsync-go/cmd/cdk/internal"
	"github.com/davemackintosh/aws-appsync-go/cmd/cdk/internal/iam"
	"github.com/davemackintosh/aws-appsync-go/cmd/cdk/internal/lambda"
	"github.com/davemackintosh/aws-appsync-go/internal/cdkutils"
	"github.com/davemackintosh/aws-appsync-go/internal/utils"
)

func NewProfileStack(app awscdk.App, api awsappsync.CfnGraphQLApi, infra *internal.InfraEntities) awscdk.Stack { //nolint:ireturn
	lambdaSources := map[string]lambda.ResolverCallback{
		"get-profile": func(dataSource awsappsync.CfnDataSource, stack awscdk.Stack, name string) {
			resolverName := name + "-resolver"
			awsappsync.NewCfnResolver(stack, &resolverName, &awsappsync.CfnResolverProps{
				ApiId:          api.AttrApiId(),
				DataSourceName: dataSource.AttrName(),
				FieldName:      utils.ToPointer("getProfile"),
				TypeName:       utils.ToPointer("Query"),
			})
		},
	}

	env := internal.InfraAccountAndRegion()
	stackName := cdkutils.NameWithEnvironment("profile")
	stack := awscdk.NewStack(app, &stackName, &awscdk.StackProps{
		Env: env,
	})

	// Create the lambdas, well assign roles next.
	lambdas := lambda.CreateLambdasFromMap(stack, api, lambdaSources, infra)

	// Give the lambdas access to the right tables and services.
	lambdas["get-profile"].AddToRolePolicy(iam.GetPolicyStatementFor(stack, iam.PolicyNameDynamoUsersRead))
	lambdas["get-profile"].AddToRolePolicy(iam.GetPolicyStatementFor(stack, iam.PolicyNameUseKMS))

	return stack
}
