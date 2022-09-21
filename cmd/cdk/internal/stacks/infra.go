package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"

	"github.com/warpspeedboilerplate/go/cmd/cdk/internal"
	"github.com/warpspeedboilerplate/go/cmd/cdk/internal/iam"
	"github.com/warpspeedboilerplate/go/internal/utils"
)

// NewInfraStack creates parts of our stack that aren't specific to any particular service
// but are needed for the whole stack, things like an VPC, SSM parameters, outputs.
func NewInfraStack(app awscdk.App) *internal.InfraEntities {
	stack := awscdk.NewStack(app, utils.ToPointer("InfraStack"), nil)

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

	// Set up the Github OIDC provider.
	iam.GithubOIDCProvider(stack)

	return &internal.InfraEntities{}
}
