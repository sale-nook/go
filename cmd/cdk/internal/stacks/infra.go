package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/davemackintosh/aws-appsync-go/cmd/cdk/internal"
	/*"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"

	"github.com/davemackintosh/aws-appsync-go/internal/cdkutils"
	"github.com/davemackintosh/aws-appsync-go/internal/utils"*/)

// NewInfraStack creates parts of our stack that aren't specific to any particular service
// but are needed for the whole stack, things like an VPC, SSM parameters.
func NewInfraStack(app awscdk.App) *internal.InfraEntities {
	/*env := internal.InfraAccountAndRegion()
	stack := awscdk.NewStack(app, utils.ToPointer(cdkutils.NameWithEnvironment("infra")), &awscdk.StackProps{
		Env: env,
	})

	// We also need a VPC to keep the lambdas safe.
	privateVpc := awsec2.NewVpc(stack, utils.ToPointer(cdkutils.NameWithEnvironment("internal")), &awsec2.VpcProps{
		SubnetConfiguration: &[]*awsec2.SubnetConfiguration{
			{
				Name:       utils.ToPointer("internal"),
				SubnetType: awsec2.SubnetType_PRIVATE_WITH_NAT, //nolint: nosnakecase
				CidrMask:   utils.ToPointer(24.0),
			},
			{
				Name:                utils.ToPointer("public"),
				SubnetType:          awsec2.SubnetType_PUBLIC, //nolint: nosnakecase
				CidrMask:            utils.ToPointer(24.0),
				MapPublicIpOnLaunch: utils.ToPointer(false),
				Reserved:            utils.ToPointer(false),
			},
		},
	})*/

	return &internal.InfraEntities{}
}
