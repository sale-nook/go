package stacks

import (
	"io/ioutil"
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2"

	"github.com/davemackintosh/cdk-appsync-go/cmd/cdk/internal"
	"github.com/davemackintosh/cdk-appsync-go/cmd/cdk/internal/iam"
	"github.com/davemackintosh/cdk-appsync-go/cmd/cdk/internal/lambda"
	"github.com/davemackintosh/cdk-appsync-go/internal/cdkutils"
	"github.com/davemackintosh/cdk-appsync-go/internal/utils"
)

func getUserMigrateFn(stack awscdk.Stack) awslambda.Function { // nolint:ireturn
	migrateFn := lambda.NewLambdaFunction("user-from-cognito-sign-up", stack, &lambda.FunctionProps{
		Entry:                 utils.ToPointer("./cmd/create-user-from-cognito"),
		RuntimeEnvironmentals: lambda.GetConfigMap(stack),
	})

	migrateFn.AddToRolePolicy(iam.GetPolicyStatementFor(stack, iam.PolicyNameDynamoUsersWrite))

	return migrateFn
}

func getWAFFirewall(stack awscdk.Stack, appsyncAPI awsappsync.CfnGraphQLApi) {
	var none interface{}

	overrideActionProperty := awswafv2.CfnWebACL_OverrideActionProperty{ // nolint: nosnakecase
		None: none,
	}
	WAF := awswafv2.NewCfnWebACL(stack, utils.ToPointer(cdkutils.NameWithStackAndEnvironment("waf-firewall", *stack.StackName())), &awswafv2.CfnWebACLProps{
		Scope: utils.ToPointer("REGIONAL"),
		DefaultAction: &awswafv2.CfnWebACL_DefaultActionProperty{ //nolint: nosnakecase
			Allow: &awswafv2.CfnWebACL_AllowActionProperty{ // nolint: nosnakecase
				CustomRequestHandling: &awswafv2.CfnWebACL_CustomRequestHandlingProperty{ //nolint: nosnakecase
					InsertHeaders: &[]awswafv2.CfnWebACL_CustomHTTPHeaderProperty{ //nolint: nosnakecase
						// Genuinly no idea why there is a requirement for this to be here but strings are needed apparently.
						{
							Name:  utils.ToPointer("cdk-appsync-go"),
							Value: utils.ToPointer("api"),
						},
					},
				},
			},
		},
		VisibilityConfig: &awswafv2.CfnWebACL_VisibilityConfigProperty{ //nolint: nosnakecase
			CloudWatchMetricsEnabled: utils.ToPointer(true),
			MetricName:               utils.ToPointer("WAFv2"),
			SampledRequestsEnabled:   utils.ToPointer(true),
		},
		Rules: &[]awswafv2.CfnWebACL_RuleProperty{ //nolint: nosnakecase
			{
				Name:     utils.ToPointer("AWS-AWSManagedRulesCommonRuleSet"),
				Priority: utils.ToPointer(1.0),
				Action: &awswafv2.CfnWebACL_BlockActionProperty{ //nolint: nosnakecase
					CustomResponse: &awswafv2.CfnWebACL_CustomResponseProperty{ //nolint: nosnakecase
						ResponseCode: utils.ToPointer(403.0),
						ResponseHeaders: &[]awswafv2.CfnWebACL_CustomHTTPHeaderProperty{ //nolint: nosnakecase
							{
								Name:  utils.ToPointer("X-Error-Code"),
								Value: utils.ToPointer("FORBIDDEN"),
							},
						},
					},
				},
				Statement: &awswafv2.CfnWebACL_StatementProperty{ //nolint: nosnakecase

					ManagedRuleGroupStatement: &awswafv2.CfnWebACL_ManagedRuleGroupStatementProperty{ //nolint: nosnakecase
						VendorName: utils.ToPointer("AWS"),
						Name:       utils.ToPointer("AWSManagedRulesCommonRuleSet"),
						ExcludedRules: &[]awswafv2.CfnWebACL_ExcludedRuleProperty{ //nolint: nosnakecase
							{
								Name: utils.ToPointer("NoUserAgent_HEADER"),
							},
						},
					},
				},
				OverrideAction: &overrideActionProperty,
				VisibilityConfig: &awswafv2.CfnWebACL_VisibilityConfigProperty{ //nolint: nosnakecase
					CloudWatchMetricsEnabled: utils.ToPointer(true),
					MetricName:               utils.ToPointer("WAFv2-AWS-AWSManagedRulesCommonRuleSet"),
					SampledRequestsEnabled:   utils.ToPointer(true),
				},
			},
			{
				Name:     utils.ToPointer("IP-based-rate-limit"),
				Priority: utils.ToPointer(2.0),
				Action: &awswafv2.CfnWebACL_BlockActionProperty{ //nolint: nosnakecase
					CustomResponse: &awswafv2.CfnWebACL_CustomResponseProperty{ //nolint: nosnakecase
						ResponseCode: utils.ToPointer(403.0),
						ResponseHeaders: &[]awswafv2.CfnWebACL_CustomHTTPHeaderProperty{ //nolint: nosnakecase
							{
								Name:  utils.ToPointer("X-Rate-Limit-Reached"),
								Value: utils.ToPointer("true"),
							},
						},
					},
				},
				Statement: &awswafv2.CfnWebACL_StatementProperty{ //nolint: nosnakecase

					RateBasedStatement: &awswafv2.CfnWebACL_RateBasedStatementProperty{ //nolint: nosnakecase
						AggregateKeyType: utils.ToPointer("IP"),
						Limit:            utils.ToPointer(100.0),
					},
				},
				OverrideAction: &overrideActionProperty,
				VisibilityConfig: &awswafv2.CfnWebACL_VisibilityConfigProperty{ //nolint: nosnakecase
					CloudWatchMetricsEnabled: utils.ToPointer(true),
					MetricName:               utils.ToPointer("WAFv2-AWS-WAF-IP-based-rate-limit"),
					SampledRequestsEnabled:   utils.ToPointer(true),
				},
			},
		},
	})

	awswafv2.NewCfnWebACLAssociation(stack, utils.ToPointer(cdkutils.NameWithStackAndEnvironment("waf-firewall-association", *stack.StackName())), &awswafv2.CfnWebACLAssociationProps{
		ResourceArn: appsyncAPI.AttrArn(),
		WebAclArn:   WAF.AttrArn(),
	})
}

func getUserIdentityPool(stack awscdk.Stack, userPool awscognito.UserPool, userPoolClient awscognito.UserPoolClient) awscognito.CfnIdentityPool { // nolint:ireturn
	identityPoolName := cdkutils.NameWithStackAndEnvironment("users-identity-pool", *stack.StackName())

	identityPool := awscognito.NewCfnIdentityPool(stack, &identityPoolName, &awscognito.CfnIdentityPoolProps{
		AllowUnauthenticatedIdentities: utils.ToPointer(true),
		CognitoIdentityProviders: &[]awscognito.CfnIdentityPool_CognitoIdentityProviderProperty{ // nolint: nosnakecase
			{
				ClientId:     userPoolClient.UserPoolClientId(),
				ProviderName: userPool.UserPoolProviderName(),
			},
		},
	})
	exportNames := internal.ExportNames()

	awscdk.NewCfnOutput(stack, exportNames.IdentityPoolID, &awscdk.CfnOutputProps{
		Value:      identityPool.Ref(),
		ExportName: exportNames.IdentityPoolID,
	})

	return identityPool
}

func getUserPoolClient(stack awscdk.Stack, userPool awscognito.UserPool) awscognito.UserPoolClient { // nolint:ireturn
	userPoolClientName := cdkutils.NameWithStackAndEnvironment("users-cdk-appsync-go-app", *stack.StackName())

	envConfing, err := internal.GetConfig()
	if err != nil {
		log.Fatalf("Failed to get CDK env config: %v", err)
	}

	userPoolClient := awscognito.NewUserPoolClient(stack, &userPoolClientName, &awscognito.UserPoolClientProps{
		UserPool: userPool,
		AuthFlows: &awscognito.AuthFlow{
			UserPassword:      utils.ToPointer(true),
			AdminUserPassword: utils.ToPointer(true),
			Custom:            utils.ToPointer(true),
			UserSrp:           utils.ToPointer(true),
		},
		GenerateSecret: utils.ToPointer(false),
		OAuth: &awscognito.OAuthSettings{
			CallbackUrls: envConfing.CognitoCallbackURLS,
			Scopes: &[]awscognito.OAuthScope{
				awscognito.OAuthScope_EMAIL(),
				awscognito.OAuthScope_OPENID(),
				awscognito.OAuthScope_PROFILE(),
			},
			Flows: &awscognito.OAuthFlows{
				AuthorizationCodeGrant: utils.ToPointer(true),
			},
		},
		SupportedIdentityProviders: &[]awscognito.UserPoolClientIdentityProvider{
			awscognito.UserPoolClientIdentityProvider_COGNITO(),
		},
	})
	exportNames := internal.ExportNames()
	awscdk.NewCfnOutput(stack, exportNames.UserPoolClientID, &awscdk.CfnOutputProps{
		Value:      userPoolClient.UserPoolClientId(),
		ExportName: exportNames.UserPoolClientID,
	})

	return userPoolClient
}

func getCognitoUserPool(stack awscdk.Stack, infra *internal.InfraEntities) awscognito.UserPool { //nolint:ireturn
	poolName := cdkutils.NameWithEnvironment("users")

	userPool := awscognito.NewUserPool(stack, &poolName, &awscognito.UserPoolProps{
		UserPoolName:      &poolName,
		SelfSignUpEnabled: utils.ToPointer(true),
		SignInAliases: &awscognito.SignInAliases{
			Email: utils.ToPointer(true),
		},
		AccountRecovery: awscognito.AccountRecovery_EMAIL_ONLY,
		RemovalPolicy:   awscdk.RemovalPolicy_DESTROY, //nolint: nosnakecase
	})

	exportNames := internal.ExportNames()
	awscdk.NewCfnOutput(stack, exportNames.UserPoolID, &awscdk.CfnOutputProps{
		Value:      userPool.UserPoolId(),
		ExportName: exportNames.UserPoolID,
	})

	return userPool
}

func getAppSyncAPI(stack awscdk.Stack, userPool awscognito.UserPool) awsappsync.CfnGraphQLApi { // nolint:ireturn
	env, err := internal.InfraAccountAndRegion()
	if err != nil {
		log.Fatalf("Failed to get CDK env config: %v", err)
	}

	content, err := ioutil.ReadFile("./web/src/graphql/schema.graphql")
	if err != nil {
		panic(err)
	}
	contentString := string(content)

	logGroup := awslogs.NewLogGroup(stack, utils.ToPointer("LogGroup"), &awslogs.LogGroupProps{
		Retention: awslogs.RetentionDays_ONE_WEEK,
	})

	appSyncName := cdkutils.NameWithEnvironment("appsync")
	appsync := awsappsync.NewCfnGraphQLApi(stack, &appSyncName, &awsappsync.CfnGraphQLApiProps{
		Name:               utils.ToPointer("cdk-appsync-go"),
		AuthenticationType: utils.ToPointer("AWS_IAM"),
		XrayEnabled:        utils.ToPointer(true),
		LogConfig: &awsappsync.CfnGraphQLApi_LogConfigProperty{
			CloudWatchLogsRoleArn: logGroup.LogGroupArn(),
			FieldLogLevel:         utils.ToPointer("ALL"),
		},
		UserPoolConfig: &awsappsync.CfnGraphQLApi_UserPoolConfigProperty{
			UserPoolId:    userPool.UserPoolId(),
			AwsRegion:     env.Region,
			DefaultAction: utils.ToPointer("ALLOW"),
		},
	})

	schemaName := cdkutils.NameWithEnvironment("appsync-schema")
	awsappsync.NewCfnGraphQLSchema(stack, &schemaName, &awsappsync.CfnGraphQLSchemaProps{
		ApiId:      appsync.AttrApiId(),
		Definition: &contentString,
	})

	exportNames := internal.ExportNames()
	awscdk.NewCfnOutput(stack, exportNames.AppSyncURL, &awscdk.CfnOutputProps{
		Value:      appsync.AttrGraphQlUrl(),
		ExportName: exportNames.AppSyncURL,
	})

	return appsync
}

func NewAppsyncStack(app awscdk.App, infra *internal.InfraEntities) (awscdk.Stack, awsappsync.CfnGraphQLApi) { //nolint:ireturn
	env, err := internal.InfraAccountAndRegion()
	if err != nil {
		log.Fatalf("Failed to get CDK env config: %v", err)
	}

	stackName := cdkutils.NameWithEnvironment("appsync")
	stack := awscdk.NewStack(app, &stackName, &awscdk.StackProps{
		Env: env,
	})

	userPool := getCognitoUserPool(stack, infra)
	userPoolClient := getUserPoolClient(stack, userPool)
	appsync := getAppSyncAPI(stack, userPool)
	identityPool := getUserIdentityPool(stack, userPool, userPoolClient)
	iam.GetAppSyncIAMRoles(stack, appsync, userPool, userPoolClient, identityPool)
	// @TODO add WAF rules to rate limit requests to the API
	// getWAFFirewall(stack, appsync)

	userPool.AddTrigger(awscognito.UserPoolOperation_POST_CONFIRMATION(), getUserMigrateFn(stack)) //nolint: nosnakecase

	return stack, appsync
}
