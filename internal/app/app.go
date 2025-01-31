package app

import (
	"fmt"
	"net/http"
	"strings"

	responseRecorder "github.com/BrunoPolaski/login-service/internal/controller/response_recorder"
	"github.com/BrunoPolaski/login-service/internal/controller/routes"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if request.Path == "" {
		return nil, fmt.Errorf("path is required")
	}

	router := routes.Init()

	httpRequest, err := http.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))
	if err != nil {
		return nil, err
	}

	for k, v := range request.Headers {
		httpRequest.Header.Set(k, v)
	}

	q := httpRequest.URL.Query()
	for k, v := range request.QueryStringParameters {
		q.Set(k, v)
	}
	httpRequest.URL.RawQuery = q.Encode()

	rr := responseRecorder.NewResponseRecorder()
	router.ServeHTTP(rr, httpRequest)

	rr.Headers["Content-Type"] = "application/json"
	rr.Headers["Access-Control-Allow-Origin"] = "*"

	return &events.APIGatewayProxyResponse{
		StatusCode: rr.StatusCode,
		Body:       rr.Body,
		Headers:    rr.Headers,
	}, nil
}
