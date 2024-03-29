package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/warpspeedboilerplate/go/cmd/create-user-from-cognito/internal/handler"
)

func main() {
	lambda.Start(handler.NewUserFromCognitoEvent)
}
