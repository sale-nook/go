package tables_test

import (
	"encoding/json"
	"testing"

	"github.com/davemackintosh/aws-appsync-go/internal/dynamo/tables"
)

const (
	testAWSAppSyncGoUser = `
	{
		"ID": "1"
	}
	`
)

func TestAWSAppSyncGoUser(t *testing.T) {
	user := tables.User{}
	err := json.Unmarshal([]byte(testAWSAppSyncGoUser), &user)
	if err != nil {
		t.Errorf("Error unmarshalling user: %v", err)
	}
	if user.ID != "1" {
		t.Errorf("Expected user ID to be 1, got %s", user.ID)
	}
}
