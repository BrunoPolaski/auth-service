package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BrunoPolaski/login-service/internal/config/logger"
	"github.com/BrunoPolaski/login-service/internal/controller/routes"
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

	if os.Getenv("ENV") == "local" {
		log.Fatal(http.ListenAndServe(":8080", r))
	} else {
		lambda.Start(httpadapter.New(r).ProxyWithContext)
	}
}
