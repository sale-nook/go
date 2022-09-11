package stacks

import (
	"fmt"
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"

	"github.com/davemackintosh/go/cmd/cdk/internal"
	"github.com/davemackintosh/go/internal/cdkutils"
	"github.com/davemackintosh/go/internal/dal"
	"github.com/davemackintosh/go/internal/utils"
)

func NewDatabaseStack(app awscdk.App, _ *internal.InfraEntities) map[string]awsdynamodb.Table {
	IDPartitionKey := utils.ToPointer("id")

	env, err := internal.InfraAccountAndRegion()
	if err != nil {
		log.Fatalf("failed to get infra account and region %s", err)
	}

	stackName := cdkutils.NameWithEnvironment("database")
	stack := awscdk.NewStack(app, &stackName, &awscdk.StackProps{
		Env: env,
	})

	kmsKey := awskms.NewKey(stack, utils.ToPointer("dynamoDBKMSKey"), &awskms.KeyProps{
		EnableKeyRotation: utils.ToPointer(true),
		Enabled:           utils.ToPointer(true),
		Policy: awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
			Statements: &[]awsiam.PolicyStatement{
				awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
					Effect: awsiam.Effect_ALLOW, // nolint: nosnakecase
					Actions: &[]*string{
						utils.ToPointer("kms:*"),
					},
					Resources: &[]*string{
						utils.ToPointer("*"),
					},
					Principals: &[]awsiam.IPrincipal{
						awsiam.NewAccountRootPrincipal(),
					},
				}),
				awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
					Effect: awsiam.Effect_ALLOW, // nolint: nosnakecase
					Actions: &[]*string{
						utils.ToPointer("kms:Encrypt"),
						utils.ToPointer("kms:Decrypt"),
					},
					Resources: &[]*string{
						utils.ToPointer("*"),
					},
					Principals: &[]awsiam.IPrincipal{
						awsiam.NewAnyPrincipal(),
					},
					Conditions: &map[string]interface{}{
						"StringEquals": &map[string]interface{}{
							"kms:CallerAccount": env.Account,
							"kms:ViaService":    utils.ToPointer(fmt.Sprintf("dynamodb.%s.amazonaws.com", *env.Region)),
						},
					},
				}),
			},
		}),
	})

	return map[string]awsdynamodb.Table{
		"users": awsdynamodb.NewTable(stack, utils.ToPointer(dal.Tables().AWSAppSyncGoUser()), &awsdynamodb.TableProps{
			PartitionKey: &awsdynamodb.Attribute{
				Name: IDPartitionKey,
				Type: awsdynamodb.AttributeType_STRING,
			},
			BillingMode:         awsdynamodb.BillingMode_PAY_PER_REQUEST,
			PointInTimeRecovery: utils.ToPointer(true),
			TableName:           utils.ToPointer(dal.Tables().AWSAppSyncGoUser()),
			EncryptionKey:       kmsKey,
			Encryption:          awsdynamodb.TableEncryption_CUSTOMER_MANAGED, //nolint: nosnakecase
		}),
	}
}
