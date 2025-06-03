package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BrunoPolaski/auth-service/internal/adapters/http/routes"
	"github.com/BrunoPolaski/auth-service/internal/config/logger"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	logger.InitLogger()
	logger.Info("Starting application")

	r := routes.Init()
	if r == nil {
		logger.Error("Failed to load routes")
		return
	}

	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		log.Fatal(http.ListenAndServe(":8080", r))
	} else {
		lambda.Start(httpadapter.New(r).ProxyWithContext)
	}
}
