package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/davemackintosh/aws-appsync-go/cmd/get-profile/handler"
)

func main() {
	lambda.Start(handler.GetProfile)
}
