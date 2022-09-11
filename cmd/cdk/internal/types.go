package internal

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"

	"github.com/davemackintosh/go/internal/utils"
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
		IdentityPoolID:   utils.ToPointer("IdentityPoolId"),
		UserPoolID:       utils.ToPointer("UserPoolId"),
		UserPoolClientID: utils.ToPointer("UserPoolClientId"),
		AppSyncURL:       utils.ToPointer("AppSyncURL"),
	}
}
