package dal_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/davemackintosh/aws-appsync-go/internal"
	"github.com/davemackintosh/aws-appsync-go/internal/dal"
)

func TestTables(t *testing.T) {
	env := os.Getenv("ENVIRONMENT")
	packageJSON := internal.GetApp()

	t.Run("withEnvironmentAndStack", func(t *testing.T) {
		t.Run("returns the correct table name", func(t *testing.T) {
			expected := fmt.Sprintf("%s-%s-%s", packageJSON.Name, env, dal.TableNamesAWSAppSyncGoUser)
			actual := dal.Tables().AWSAppSyncGoUser()

			if actual != expected {
				t.Errorf("expected %s, got %s", expected, actual)
			}
		})
	})
	t.Run("withARN", func(t *testing.T) {
		t.Run("returns the correct table name", func(t *testing.T) {
			expected := fmt.Sprintf("arn:aws:dynamodb:*:*:table/%s-%s-%s", packageJSON.Name, env, dal.TableNamesAWSAppSyncGoUser)
			actual := dal.TableARNs().AWSAppSyncGoUser()

			if actual != expected {
				t.Errorf("expected %s, got %s", expected, actual)
			}
		})
	})
}
