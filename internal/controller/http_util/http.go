package http_util

import (
	"io"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
)

func convertHeaders(headers *http.Request) map[string]string {
	headersMap := make(map[string]string)
	for k, v := range headers.Header {
		headersMap[k] = v[0]
	}
	return headersMap
}

func convertQueryStringParameters(queryStringParameters url.Values) map[string]string {
	queryStringParams := make(map[string]string)
	for k, v := range queryStringParameters {
		queryStringParams[k] = v[0]
	}
	return queryStringParams
}

func convertBodyToString(body *io.ReadCloser) (string, error) {
	reqBody, err := io.ReadAll(*body)
	if err != nil {
		return "", err
	}
	return string(reqBody), nil
}

func ConvertRequestToAPIGatewayProxyRequest(r *http.Request) *events.APIGatewayProxyRequest {
	headers := convertHeaders(r)
	queryStringParams := convertQueryStringParameters(r.URL.Query())
	body, err := convertBodyToString(&r.Body)
	if err != nil {
		return nil
	}

	return &events.APIGatewayProxyRequest{
		HTTPMethod:            r.Method,
		Body:                  body,
		Headers:               headers,
		QueryStringParameters: queryStringParams,
		Path:                  r.URL.Path,
	}
}
