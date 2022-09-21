package iam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/davemackintosh/go/internal/utils"
)

func GetUserPoolGroups() map[string]*awscognito.CfnUserPoolGroupProps {
	return map[string]*awscognito.CfnUserPoolGroupProps{
		"registered": {
			// This group is used for users who have not yet subscribed to
			// a paid plan. They are not allowed to access any of the
			// premium content.
			GroupName: utils.ToPointer("registered"),
		},
		"subscribed": {
			// This group is used for users who have subscribed to a paid
			// plan. They are allowed to access all content.
			GroupName: utils.ToPointer("subscribed"),
		},
	}
}
