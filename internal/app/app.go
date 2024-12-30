package app

import (
	"net/http"
	"strings"

	responseRecorder "github.com/BrunoPolaski/login-service/internal/controller/response_recorder"
	"github.com/BrunoPolaski/login-service/internal/controller/routes"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	if request.Path == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Path cannot be empty",
		}
	}

	router := routes.Init()

	httpRequest, _ := http.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))

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
	}
}
