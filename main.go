package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BrunoPolaski/login-service/internal/app"
	"github.com/BrunoPolaski/login-service/internal/config/logger"
	"github.com/BrunoPolaski/login-service/internal/controller/http_util"
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

	if os.Getenv("AWS_LAMBDA_RUNTIME_API") == "" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			request := http_util.ConvertRequestToAPIGatewayProxyRequest(r)

			response, err := app.Handler(request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			for k, v := range response.Headers {
				w.Header().Set(k, v)
			}
			w.WriteHeader(response.StatusCode)
			w.Write([]byte(response.Body))
		})

		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		lambda.Start(app.Handler)
	}
}
