package lambda

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"

	"github.com/warpspeedboilerplate/go/cmd/cdk/internal"
	"github.com/warpspeedboilerplate/go/cmd/cdk/internal/iam"
	binternal "github.com/warpspeedboilerplate/go/internal"
	"github.com/warpspeedboilerplate/go/internal/utils"
)

type FunctionProps struct {
	RuntimeEnvironmentals *map[string]*string
	Entry                 *string
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
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		panic("ENVIRONMENT environment variable is not set")
	}

	return &map[string]*string{
		"ENVIRONMENT": utils.ToPointer(environment),
	}
}

func NewLambdaFunction(name string, stack awscdk.Stack, props *FunctionProps) awslambda.Function { //nolint:ireturn
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		panic("ENVIRONMENT environment variable is not set")
	}

	packageJSON := binternal.GetApp()

	buildHooks := BuildHooks{
		GithubAccessToken: awsssm.StringParameter_ValueFromLookup(stack, utils.ToPointer(fmt.Sprintf("/%s/%s/GITHUB_ACCESS_TOKEN", packageJSON.Name, environment))), //nolint: nosnakecase
	}

	funcProps := awscdklambdagoalpha.GoFunctionProps{
		Timeout: awscdk.Duration_Seconds(utils.ToPointer(10.0)), // nolint: nosnakecase
		Entry:   props.Entry,
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

type DataResolverTypes string

var (
	DataResolverTypesQUERY        = DataResolverTypes("Query")        //nolint: gochecknoglobals
	DataResolverTypesMUTATION     = DataResolverTypes("Mutation")     //nolint: gochecknoglobals
	DataResolverTypesSUBSCRIPTION = DataResolverTypes("Subscription") //nolint: gochecknoglobals
)

type DataResolverProps struct {
	CmdPath       string
	FieldName     string
	TypeName      DataResolverTypes
	FunctionProps FunctionProps
	IAMPolicies   []awsiam.PolicyStatement
}

func CreateLambdasFromMap(stack awscdk.Stack, api awsappsync.CfnGraphQLApi, lambdaSources []DataResolverProps, infra *internal.InfraEntities) map[string]awslambda.Function {
	lambdaFunctions := make(map[string]awslambda.Function)

	for _, lambdaConfig := range lambdaSources {
		lambdaSource := lambdaConfig.CmdPath
		newLambda := NewLambdaFunction(lambdaSource+"-fn", stack, &FunctionProps{
			Entry:                 utils.ToPointer("./cmd/" + lambdaSource),
			RuntimeEnvironmentals: GetConfigMap(stack),
			API:                   &api,
		})

		if lambdaConfig.IAMPolicies != nil {
			for _, policy := range lambdaConfig.IAMPolicies {
				newLambda.AddToRolePolicy(policy)
			}
		}

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

		resolverName := strings.ReplaceAll(lambdaSource+"-resolver", "-", "_")
		awsappsync.NewCfnResolver(stack, &resolverName, &awsappsync.CfnResolverProps{
			ApiId:          api.AttrApiId(),
			DataSourceName: ds.AttrName(),
			FieldName:      &lambdaConfig.FieldName,
			TypeName:       utils.ToPointer(string(lambdaConfig.TypeName)),
		})

		// Push lambda function to our map.
		lambdaFunctions[lambdaSource] = newLambda
	}

	return lambdaFunctions
}
