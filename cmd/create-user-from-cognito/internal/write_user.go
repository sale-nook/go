package internal

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/warpspeedboilerplate/go/internal/dal"
	"github.com/warpspeedboilerplate/go/internal/dynamo/tables"
	"github.com/warpspeedboilerplate/go/internal/utils"
)

func WriteUserToTableWithClient(ctx context.Context, client *dynamodb.Client, user tables.User) error {
	record, err := attributevalue.MarshalMap(user)
	if err != nil {
		return fmt.Errorf("error marshalling user: %w", err)
	}

	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: utils.ToPointer(dal.Tables().AWSAppSyncGoUser()),
		Item:      record,
	})
	if err != nil {
		return fmt.Errorf("failed to insert new user %w", err)
	}

	return nil
}
