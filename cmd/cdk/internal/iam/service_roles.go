package iam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"

	"github.com/davemackintosh/cdk-appsync-go/internal/utils"
)

type ServiceRoles string

const (
	ServiceRoleAppsyncLambda = "appsync-lambda"
)

func GetServiceRoles(stack awscdk.Stack) map[ServiceRoles]awsiam.Role {
	appsyncLambdaRole := awsiam.NewRole(stack, utils.ToPointer(ServiceRoleAppsyncLambda), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(utils.ToPointer("appsync"), nil),
	})

	appsyncLambdaRole.AddManagedPolicy(awsiam.ManagedPolicy_FromAwsManagedPolicyName(utils.ToPointer("AWSLambda_FullAccess")))

	return map[ServiceRoles]awsiam.Role{
		ServiceRoleAppsyncLambda: appsyncLambdaRole,
	}
}

var GetServiceRoleFor utils.MemoizedCB[ServiceRoles, awsiam.Role, awscdk.Stack] = utils.Memoize( //nolint: gochecknoglobals
	func(key ServiceRoles, stack awscdk.Stack) awsiam.Role {
		role, ok := GetServiceRoles(stack)[key]

		if !ok {
			return nil
		}

		return role
	},
)
