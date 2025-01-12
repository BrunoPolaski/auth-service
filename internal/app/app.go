package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/BrunoPolaski/login-service/internal/config/logger"
	responseRecorder "github.com/BrunoPolaski/login-service/internal/controller/response_recorder"
	"github.com/BrunoPolaski/login-service/internal/controller/routes"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger.Info(
		fmt.Sprintf("Request path: %s \nRequest HTTP method: %s \nRequest body: %s \nRequest headers: %v \nRequest query string parameters: %v \n", request.Path, request.HTTPMethod, request.Body, request.Headers, request.QueryStringParameters),
	)
	if request.Path == "" {
		logger.Error("Path cannot be empty")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Path cannot be empty",
		}, fmt.Errorf("Path cannot be empty")
	}

	router := routes.Init()

	httpRequest, err := http.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))
	if err != nil {
		logger.Error(
			fmt.Sprintf("Error creating HTTP request: %s", err),
		)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error creating HTTP request",
		}, fmt.Errorf("Error creating HTTP request")
	}

	for k, v := range request.Headers {
		httpRequest.Header.Add(k, v)
	}

	q := httpRequest.URL.Query()
	for k, v := range request.QueryStringParameters {
		q.Add(k, v)
	}
	httpRequest.URL.RawQuery = q.Encode()

	rr := responseRecorder.NewResponseRecorder()
	router.ServeHTTP(rr, httpRequest)

	return events.APIGatewayProxyResponse{
		StatusCode: rr.StatusCode,
		Body:       rr.Body,
		Headers:    rr.Headers,
	}, nil
}
