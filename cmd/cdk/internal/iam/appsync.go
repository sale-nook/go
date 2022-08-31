package iam

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"

	"github.com/davemackintosh/cdk-appsync-go/internal/cdkutils"
	"github.com/davemackintosh/cdk-appsync-go/internal/utils"
)

func GetAppSyncIAMRoles(stack awscdk.Stack, api awsappsync.CfnGraphQLApi, userPool awscognito.UserPool, userPoolclient awscognito.UserPoolClient, identityPool awscognito.CfnIdentityPool) awscognito.CfnIdentityPoolRoleAttachment { //nolint: ireturn
	apiResourceBase := fmt.Sprintf("arn:aws:appsync:%s:%s:apis/%s", *stack.Region(), *stack.Account(), *api.AttrApiId())
	unauthenticatedUserRole := awsiam.NewRole(stack, utils.ToPointer(cdkutils.NameWithStackAndEnvironment("appsync-guest", *stack.StackName())), &awsiam.RoleProps{
		AssumedBy: awsiam.NewFederatedPrincipal(utils.ToPointer("cognito-identity.amazonaws.com"), &map[string]interface{}{
			"StringEquals": map[string]interface{}{
				"cognito-identity.amazonaws.com:aud": identityPool.Ref(),
			},
			"ForAnyValue:StringLike": map[string]interface{}{
				"cognito-identity.amazonaws.com:amr": "unauthenticated",
			},
		}, utils.ToPointer("sts:AssumeRoleWithWebIdentity")),
		InlinePolicies: &map[string]awsiam.PolicyDocument{
			"guest-appsync": awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
				Statements: &[]awsiam.PolicyStatement{
					awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
						Effect: awsiam.Effect_ALLOW, // nolint: nosnakecase
						Actions: &[]*string{
							utils.ToPointer("appsync:GraphQL"),
							utils.ToPointer("cognito-sync:*"),
							utils.ToPointer("cognito-identity:*"),
						},
						Resources: &[]*string{
							utils.ToPointer(fmt.Sprintf("%s/types/mutation/fields/register", apiResourceBase)),
							utils.ToPointer(fmt.Sprintf("%s/types/mutation/fields/login", apiResourceBase)),
						},
					}),
				},
			}),
		},
	})

	authenticatedUserRole := awsiam.NewRole(stack, utils.ToPointer(cdkutils.NameWithStackAndEnvironment("appsync-user", *stack.StackName())), &awsiam.RoleProps{
		AssumedBy: awsiam.NewFederatedPrincipal(utils.ToPointer("cognito-identity.amazonaws.com"), &map[string]interface{}{
			"StringEquals": map[string]interface{}{
				"cognito-identity.amazonaws.com:aud": identityPool.Ref(),
			},
			"ForAnyValue:StringLike": map[string]interface{}{
				"cognito-identity.amazonaws.com:amr": "authenticated",
			},
		}, utils.ToPointer("sts:AssumeRoleWithWebIdentity")),
		ManagedPolicies: &[]awsiam.IManagedPolicy{
			awsiam.ManagedPolicy_FromAwsManagedPolicyName(utils.ToPointer("service-role/AWSLambdaBasicExecutionRole")), // nolint: nosnakecase
		},
		InlinePolicies: &map[string]awsiam.PolicyDocument{
			"user-appsync": awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
				Statements: &[]awsiam.PolicyStatement{
					awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
						Effect: awsiam.Effect_ALLOW, // nolint: nosnakecase
						Actions: &[]*string{
							utils.ToPointer("appsync:GraphQL"),
							utils.ToPointer("cognito-sync:*"),
							utils.ToPointer("cognito-identity:*"),
						},
						Resources: &[]*string{
							utils.ToPointer(fmt.Sprintf("%s/*", apiResourceBase)),
						},
					}),
				},
			}),
		},
	})

	wafAssoc := awscognito.NewCfnIdentityPoolRoleAttachment(stack, utils.ToPointer(cdkutils.NameWithStackAndEnvironment("appsync-identity-pool-role-attachment", *stack.StackName())), &awscognito.CfnIdentityPoolRoleAttachmentProps{
		IdentityPoolId: identityPool.Ref(),
		Roles: map[string]interface{}{
			"authenticated":   authenticatedUserRole.RoleArn(),
			"unauthenticated": unauthenticatedUserRole.RoleArn(),
		},
		RoleMappings: &map[string]awscognito.CfnIdentityPoolRoleAttachment_RoleMappingProperty{ // nolint: nosnakecase
			"authenticated": {
				AmbiguousRoleResolution: utils.ToPointer("AuthenticatedRole"),
				IdentityProvider:        utils.ToPointer(fmt.Sprintf("cognito-idp.%s.amazonaws.com/%s:%s", *stack.Region(), *userPool.UserPoolId(), *userPoolclient.UserPoolClientId())),
				Type:                    utils.ToPointer("Rules"),
				RulesConfiguration: &awscognito.CfnIdentityPoolRoleAttachment_RulesConfigurationTypeProperty{ // nolint: nosnakecase
					Rules: &[]awscognito.CfnIdentityPoolRoleAttachment_MappingRuleProperty{ // nolint: nosnakecase
						{
							Claim:     utils.ToPointer("sub"),
							MatchType: utils.ToPointer("Equals"),
							RoleArn:   authenticatedUserRole.RoleArn(),
							Value:     utils.ToPointer("authed"),
						},
					},
				},
			},
		},
	})

	wafAssoc.AddDependsOn(api)

	return wafAssoc
}
