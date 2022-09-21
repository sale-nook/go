package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/warpspeedboilerplate/go/cmd/get-profile/handler"
)

func main() {
	lambda.Start(handler.GetProfile)
}
