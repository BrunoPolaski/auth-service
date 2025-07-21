package main

import (
	"log"
	"net/http"

	"github.com/BrunoPolaski/auth-service/internal/adapters/http/routes"
	"github.com/BrunoPolaski/auth-service/internal/infra/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()
	logger.Info("Starting application")

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error loading env file: " + err.Error())
	}

	r := routes.Init()
	if r == nil {
		logger.Error("Failed to load routes")
		return
	}

	log.Fatal(http.ListenAndServe(":8080", r))
}
