package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/davemackintosh/aws-appsync-go/cmd/cdk/internal"
	"github.com/davemackintosh/aws-appsync-go/internal/cdkutils"
	"github.com/davemackintosh/aws-appsync-go/internal/utils"
)

// NewInfraStack creates parts of our stack that aren't specific to any particular service
// but are needed for the whole stack, things like an VPC, SSM parameters, outputs.
func NewInfraStack(app awscdk.App) *internal.InfraEntities {
	env := internal.InfraAccountAndRegion()
	stack := awscdk.NewStack(app, utils.ToPointer(cdkutils.NameWithEnvironment("infra")), &awscdk.StackProps{
		Env: env,
	})

	// No need for public subnets in our application
	// so this is commented out but if you want to use it
	// uncomment the following lines.
	/*vpc := awsec2.NewVpc(stack, utils.ToPointer(cdkutils.NameWithEnvironment("internal")), &awsec2.VpcProps{
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

	// Write the region to the outputs.
	awscdk.NewCfnOutput(stack, utils.ToPointer(cdkutils.NameWithEnvironment("region")), &awscdk.CfnOutputProps{
		Value: env.Region,
	})

	return &internal.InfraEntities{}
}
