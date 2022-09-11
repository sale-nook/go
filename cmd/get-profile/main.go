package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/davemackintosh/go/cmd/get-profile/handler"
)

func main() {
	lambda.Start(handler.GetProfile)
}
