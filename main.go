package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BrunoPolaski/login-service/internal/app"
	"github.com/BrunoPolaski/login-service/internal/config/logger"
	"github.com/BrunoPolaski/login-service/internal/controller/routes"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file. Error: %s", err)
	}
	logger.InitLogger()
	logger.Info("Starting application")

	if os.Getenv("ENV") == "local" {
		r := routes.Init()

		log.Fatal(http.ListenAndServe(":8080", r))
	}

	lambda.Start(app.Handler)
}
