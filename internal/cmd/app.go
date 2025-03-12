package app

import (
	"fmt"
	"net/http"
	"strings"

	internalHttp "github.com/BrunoPolaski/login-service/internal/interfaces/http"
	"github.com/BrunoPolaski/login-service/internal/interfaces/http/routes"
	"github.com/aws/aws-lambda-go/events"
)

var router http.Handler

func init() {
	router = routes.Init()
}

func Handler(request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if request.Path == "" {
		return nil, fmt.Errorf("path is required")
	}

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

	rr := internalHttp.NewResponseRecorder()
	router.ServeHTTP(rr, httpRequest)

	rr.Headers["Content-Type"] = "application/json"
	rr.Headers["Access-Control-Allow-Origin"] = "*"

	return &events.APIGatewayProxyResponse{
		StatusCode: rr.StatusCode,
		Body:       rr.Body,
		Headers:    rr.Headers,
	}, nil
}
