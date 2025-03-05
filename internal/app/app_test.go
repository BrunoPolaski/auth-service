package app_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/BrunoPolaski/login-service/internal/app"
	"github.com/BrunoPolaski/login-service/internal/tests"
	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Setenv("ENV", "dev")
	t.Setenv("LOG_LEVEL", "info")
	t.Run("should add query string parameters to request", func(t *testing.T) {
		request := &events.APIGatewayProxyRequest{
			Path: "/health",
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

		tests.AssertEqual(t, "key=value", httpRequest.URL.RawQuery)

		response, err := app.Handler(request)

		tests.AssertNil(t, err)
		tests.AssertEqual(t, http.StatusOK, response.StatusCode)
	})

	t.Run("should add request headers", func(t *testing.T) {
		request := &events.APIGatewayProxyRequest{
			Path: "/health",
			Headers: map[string]string{
				"key": "value",
			},
		}

		httpRequest, _ := http.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))

		for k, v := range request.Headers {
			httpRequest.Header.Add(k, v)
		}

		tests.AssertEqual(t, "value", httpRequest.Header.Get("key"))

		response, err := app.Handler(request)

		tests.AssertNil(t, err)
		tests.AssertEqual(t, http.StatusOK, response.StatusCode)
	})

	t.Run("should return response", func(t *testing.T) {
		request := &events.APIGatewayProxyRequest{
			Path: "/health",
		}

		response, err := app.Handler(request)

		tests.AssertNil(t, err)
		tests.AssertEqual(t, http.StatusOK, response.StatusCode)
	})

	t.Run("should return error when passing invalid request", func(t *testing.T) {
		request := &events.APIGatewayProxyRequest{}

		response, err := app.Handler(request)

		tests.AssertEqual(t, http.StatusBadRequest, err.Code)
		tests.AssertEqual(t, "Path is required", err.Message)
		tests.AssertNotNil(t, err)
		tests.AssertNil(t, response)
	})
}
