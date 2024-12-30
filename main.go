package main

import (
	"log"

	"github.com/BrunoPolaski/login-service/internal/app"
	"github.com/BrunoPolaski/login-service/internal/config/logger"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file. Error: %s", err)
	}
	logger.InitLogger()

	lambda.Start(app.Handler)
}
