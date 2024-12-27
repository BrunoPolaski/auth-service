package main

import (
	"github.com/BrunoPolaski/login-service/internal/app"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(app.Handler)
}
