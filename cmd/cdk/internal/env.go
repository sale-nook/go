package internal

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func requireAccountID() *string {
	accountID := os.Getenv("AWS_ACCOUNT_ID")

	if accountID == "" {
		panic("AWS_ACCOUNT_ID is not set")
	}

	return &accountID
}

func requireRegion() *string {
	region := os.Getenv("AWS_REGION")

	if region == "" {
		panic("AWS_REGION is not set")
	}

	return &region
}

func InfraAccountAndRegion() *awscdk.Environment {
	return &awscdk.Environment{
		Account: requireAccountID(),
		Region:  requireRegion(),
	}
}
