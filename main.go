package main

import (
	"log"
	"net/http"

	"github.com/BrunoPolaski/auth-service/internal/adapters/http/routes"
	"github.com/BrunoPolaski/auth-service/internal/config/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			log.Fatalf("%s", err)
		}
	}
	logger.InitLogger()
	logger.Info("Starting application")

	r := routes.Init()
	if r == nil {
		logger.Error("Failed to load routes")
		return
	}

	log.Fatal(http.ListenAndServe(":8080", r))
}
