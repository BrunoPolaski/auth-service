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
	router := routes.Init()
	httpRequest, err := http.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create HTTP request: %v", err))
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, fmt.Errorf("failed to create HTTP request: %w", err)
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
