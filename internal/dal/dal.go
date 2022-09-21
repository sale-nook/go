package dal

import "github.com/warpspeedboilerplate/go/internal/cdkutils"

const (
	TableNamesAWSAppSyncGoUser string = "go_user"
	stackName                  string = "database"
)

type (
	TableName func() string
)

type TableNames struct {
	AWSAppSyncGoUser TableName
}

func withARN(tableName TableName) TableName {
	return func() string {
		return "arn:aws:dynamodb:*:*:table/" + tableName()
	}
}

func withEnvironment(tableName string) TableName {
	return func() string {
		return cdkutils.NameWithEnvironment(tableName)
	}
}

func Tables() TableNames {
	return TableNames{
		AWSAppSyncGoUser: withEnvironment(TableNamesAWSAppSyncGoUser),
	}
}

func TableARNs() TableNames {
	return TableNames{
		AWSAppSyncGoUser: withARN(
			withEnvironment(TableNamesAWSAppSyncGoUser),
		),
	}
}
