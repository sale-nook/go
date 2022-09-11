package iam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/davemackintosh/cdk-appsync-go/internal/utils"
)

func GithubOIDCProvider(stack awscdk.Stack) {
	provider := awsiam.NewOpenIdConnectProvider(
		stack,
		utils.ToPointer("GithubOIDCProvider"),
		&awsiam.OpenIdConnectProviderProps{
			Url: utils.ToPointer("https://token.actions.githubusercontent.com"),
			ClientIds: &[]*string{
				utils.ToPointer("sts.amazonaws.com"),
			},
			Thumbprints: &[]*string{
				utils.ToPointer("6938fd4d98bab03faadb97b34396831e3780aea1"),
			},
		},
	)

	role := awsiam.NewRole(
		stack,
		utils.ToPointer("GithubActionsRole"),
		&awsiam.RoleProps{
			AssumedBy: awsiam.NewWebIdentityPrincipal(
				provider.OpenIdConnectProviderArn(),
				&map[string]interface{}{
					"StringLike": &map[string]interface{}{
						"token.actions.githubusercontent.com:sub": "repo:davemackintosh/cdk-appsync-go:*",
					},
				},
			),
			ManagedPolicies: &[]awsiam.IManagedPolicy{
				awsiam.ManagedPolicy_FromAwsManagedPolicyName(
					utils.ToPointer("AdministratorAccess"),
				),
			},
			RoleName:           utils.ToPointer("GithubActionsRole"),
			MaxSessionDuration: awscdk.Duration_Hours(utils.ToPointer(1.0)),
		},
	)

	awscdk.NewCfnOutput(
		stack,
		utils.ToPointer("GithubActionsRoleArn"),
		&awscdk.CfnOutputProps{
			Value: role.RoleArn(),
		},
	)
}
