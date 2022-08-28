package lambda

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"

	"github.com/davemackintosh/aws-appsync-go/cmd/cdk/internal"
	"github.com/davemackintosh/aws-appsync-go/cmd/cdk/internal/iam"
	binternal "github.com/davemackintosh/aws-appsync-go/internal"
	"github.com/davemackintosh/aws-appsync-go/internal/utils"
)

type FunctionProps struct {
	BuildEnvironmentals   *map[string]*string
	RuntimeEnvironmentals *map[string]*string
	Source                *string
	API                   *awsappsync.CfnGraphQLApi
	Vpc                   *awsec2.Vpc
	URLProps              *awslambda.FunctionUrlOptions
}

type BuildHooks struct {
	GithubAccessToken *string
}

func (b *BuildHooks) BeforeBundling(inDir *string, _ *string) *[]*string {
	return &[]*string{
		utils.ToPointer(
			fmt.Sprintf("echo \"machine github.com login %s\" >> ~/.netrc && chmod 0600 ~/.netrc", *b.GithubAccessToken),
		),
	}
}

func (b *BuildHooks) AfterBundling(inputDir *string, outputDir *string) *[]*string {
	return &[]*string{
		utils.ToPointer("rm ~/.netrc"),
	}
}

// GetConfigMap returns a map of config values, typycally used for
// setting the.
func GetConfigMap(stack awscdk.Stack) *map[string]*string {
	packageJSON := binternal.GetApp()

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		panic("ENVIRONMENT environment variable is not set")
	}

	oauthCallbackBaseURL := os.Getenv("OAUTH_CALLBACK_ROOT")
	if oauthCallbackBaseURL == "" {
		panic("OAUTH_CALLBACK_ROOT environment variable is not set")
	}

	exportNames := internal.ExportNames()

	return &map[string]*string{
		"ENVIRONMENT":         utils.ToPointer(environment),
		"OAUTH_CALLBACK_ROOT": utils.ToPointer(oauthCallbackBaseURL),
		"USER_POOL_CLIENT_ID": awscdk.Fn_ImportValue(exportNames.UserPoolClientID),                                                                                      //nolint: nosnakecase
		"USER_POOL_ID":        awscdk.Fn_ImportValue(exportNames.UserPoolID),                                                                                            //nolint: nosnakecase
		"GITHUB_ACCESS_TOKEN": awsssm.StringParameter_ValueFromLookup(stack, utils.ToPointer(fmt.Sprintf("/%s/%s/GITHUB_ACCESS_TOKEN", packageJSON.Name, environment))), //nolint: nosnakecase
	}
}

func NewLambdaFunction(name string, stack awscdk.Stack, props *FunctionProps) awslambda.Function { //nolint:ireturn
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		panic("ENVIRONMENT environment variable is not set")
	}

	buildHooks := BuildHooks{
		GithubAccessToken: awsssm.StringParameter_ValueFromLookup(stack, utils.ToPointer(fmt.Sprintf("/%s/GITHUB_ACCESS_TOKEN", environment))), //nolint: nosnakecase
	}

	funcProps := awscdklambdagoalpha.GoFunctionProps{
		Timeout: awscdk.Duration_Seconds(utils.ToPointer(10.0)), // nolint: nosnakecase
		Entry:   props.Source,
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			CommandHooks: &buildHooks,
			Environment:  props.RuntimeEnvironmentals,
			GoBuildFlags: &[]*string{
				utils.ToPointer("-ldflags=\"-X 'main.CommitID=$(git rev-parse HEAD) -s -w'\""),
			},
		},
		Environment:                  props.RuntimeEnvironmentals,
		DeadLetterQueueEnabled:       utils.ToPointer(true),
		ReservedConcurrentExecutions: utils.ToPointer(10.0),
		LogRetention:                 awslogs.RetentionDays_ONE_WEEK,
	}

	if props.Vpc != nil {
		funcProps.Vpc = *props.Vpc
	}

	newFunction := awscdklambdagoalpha.NewGoFunction(stack, utils.ToPointer(name), &funcProps)

	if props.URLProps != nil {
		url := newFunction.AddFunctionUrl(props.URLProps)

		awscdk.NewCfnOutput(stack, utils.ToPointer(fmt.Sprintf("%s-url", name)), &awscdk.CfnOutputProps{
			Value: url.Url(),
		})
	}

	return newFunction
}

type ResolverCallback = func(awsappsync.CfnDataSource, awscdk.Stack, string)

func CreateLambdasFromMap(stack awscdk.Stack, api awsappsync.CfnGraphQLApi, lambdaSources map[string]ResolverCallback, infra *internal.InfraEntities) map[string]awslambda.Function {
	lambdaFunctions := make(map[string]awslambda.Function)

	for lambdaSource, resolverCallback := range lambdaSources {
		newLambda := NewLambdaFunction(lambdaSource+"-fn", stack, &FunctionProps{
			Source:                utils.ToPointer("./cmd/" + lambdaSource),
			RuntimeEnvironmentals: GetConfigMap(stack),
			API:                   &api,
		})

		dataSourceName := strings.ReplaceAll(lambdaSource+"-data-source", "-", "_")
		ds := awsappsync.NewCfnDataSource(stack, &dataSourceName, &awsappsync.CfnDataSourceProps{
			ApiId: api.AttrApiId(),
			Name:  &dataSourceName,
			Type:  utils.ToPointer("AWS_LAMBDA"),
			LambdaConfig: &awsappsync.CfnDataSource_LambdaConfigProperty{ //nolint: nosnakecase
				LambdaFunctionArn: newLambda.FunctionArn(),
			},
			ServiceRoleArn: iam.GetServiceRoleFor(iam.ServiceRoleAppsyncLambda, stack).RoleArn(),
		})

		lambdaFunctions[lambdaSource] = newLambda

		resolverCallback(ds, stack, lambdaSource)
	}

	return lambdaFunctions
}
