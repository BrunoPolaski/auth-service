package main

import (
	"log"
	"net/http"
	"os"

	app "github.com/BrunoPolaski/login-service/internal/cmd"
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

		server := &http.Server{
			Addr:    ":8080",
			Handler: r,
		}

		log.Fatal(server.ListenAndServe())
	}

	lambda.Start(app.Handler)
}
