package iam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"

	"github.com/davemackintosh/cdk-appsync-go/internal/dal"
	"github.com/davemackintosh/cdk-appsync-go/internal/utils"
)

type PolicyNamesRoles string

const (
	PolicyNameDynamoUsersWrite = "dynamo-users-write"
	PolicyNameDynamoUsersRead  = "dynamo-users-read"
	PolicyNameUseKMS           = "dynamo-kms-decrypt"
)

func GetPolicyStatements(stack awscdk.Stack) map[PolicyNamesRoles]awsiam.PolicyStatement {
	return map[PolicyNamesRoles]awsiam.PolicyStatement{
		PolicyNameDynamoUsersWrite: awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
			Effect: awsiam.Effect_ALLOW, // nolint: nosnakecase
			Actions: &[]*string{
				utils.ToPointer("dynamodb:PutItem"),
			},
			Resources: &[]*string{
				utils.ToPointer(dal.TableARNs().AWSAppSyncGoUser()),
			},
		}),
		PolicyNameDynamoUsersRead: awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
			Effect: awsiam.Effect_ALLOW, // nolint: nosnakecase
			Actions: &[]*string{
				utils.ToPointer("dynamodb:GetItem"),
			},
			Resources: &[]*string{
				utils.ToPointer(dal.TableARNs().AWSAppSyncGoUser()),
			},
		}),
		PolicyNameUseKMS: awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
			Effect: awsiam.Effect_ALLOW, // nolint: nosnakecase
			Actions: &[]*string{
				utils.ToPointer("kms:*"),
			},
			Resources: &[]*string{
				utils.ToPointer("arn:aws:dynamodb:*:*:table/*"),
			},
		}),
	}
}

func GetPolicyStatementFor(stack awscdk.Stack, policyName PolicyNamesRoles) awsiam.PolicyStatement { // nolint: ireturn
	statement, ok := GetPolicyStatements(stack)[policyName]

	if !ok {
		return nil
	}

	return statement
}
