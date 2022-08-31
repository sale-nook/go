package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davemackintosh/cdk-appsync-go/cmd/create-user-from-cognito/internal/handler"
)

func main() {
	lambda.Start(handler.NewUserFromCognitoEvent)
}
