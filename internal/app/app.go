package app

import (
	"net/http"
	"strings"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	responseRecorder "github.com/BrunoPolaski/login-service/internal/controller/response_recorder"
	"github.com/BrunoPolaski/login-service/internal/controller/routes"
	"github.com/aws/aws-lambda-go/events"
)

var router http.Handler

func init() {
	router = routes.Init()
}

func Handler(request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, *rest_err.RestErr) {
	if request == nil || request.Path == "" {
		return nil, rest_err.NewBadRequestError("Path is required")
	}

	httpRequest, err := http.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
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
