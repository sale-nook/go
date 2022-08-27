package internal

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"

	"github.com/davemackintosh/aws-appsync-go/internal/utils"
)

type InfraEntities struct {
	UserPool awscognito.UserPool
}

type ExportName struct {
	IdentityPoolID   *string
	UserPoolID       *string
	UserPoolClientID *string
	AppSyncURL       *string
}

func ExportNames() ExportName {
	return ExportName{
		IdentityPoolID:   utils.ToPointer("aws-appsync-goIdentityPoolId"),
		UserPoolID:       utils.ToPointer("aws-appsync-goUserPoolId"),
		UserPoolClientID: utils.ToPointer("aws-appsync-goUserPoolClientId"),
		AppSyncURL:       utils.ToPointer("aws-appsync-goAppSyncURL"),
	}
}
