package app_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/BrunoPolaski/login-service/internal/app"
	"github.com/aws/aws-lambda-go/events"
	"github.com/joho/godotenv"
)

func TestHandler(t *testing.T) {
	godotenv.Load("../../.env")
	t.Run("should add query string parameters to request", func(t *testing.T) {
		request := events.APIGatewayProxyRequest{
			Path: "/",
			QueryStringParameters: map[string]string{
				"key": "value",
			},
		}

		httpRequest, _ := http.NewRequest(request.HTTPMethod, "", nil)

		q := httpRequest.URL.Query()
		for k, v := range request.QueryStringParameters {
			q.Add(k, v)
		}
		httpRequest.URL.RawQuery = q.Encode()

		if httpRequest.URL.RawQuery != "key=value" {
			t.Errorf("Expected query string to be key=value, got %s", httpRequest.URL.RawQuery)
		}

		response := app.Handler(request)

		if response.StatusCode != 200 {
			t.Errorf("Expected status code to be 200, got %d", response.StatusCode)
		}
	})

	t.Run("should add request headers", func(t *testing.T) {
		request := events.APIGatewayProxyRequest{
			Path: "/",
			Headers: map[string]string{
				"key": "value",
			},
		}

		httpRequest, _ := http.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))

		for k, v := range request.Headers {
			httpRequest.Header.Add(k, v)
		}

		if httpRequest.Header.Get("key") != "value" {
			t.Errorf("Expected header key to be value, got %s", httpRequest.Header.Get("key"))
		}

		response := app.Handler(request)

		if response.StatusCode != 200 {
			t.Errorf("Expected status code to be 200, got %d", response.StatusCode)
		}
	})

	t.Run("should return response", func(t *testing.T) {
		request := events.APIGatewayProxyRequest{
			Path: "/",
		}

		response := app.Handler(request)

		if response.StatusCode != 200 {
			t.Errorf("Expected status code to be 200, got %d", response.StatusCode)
		}
	})
}
