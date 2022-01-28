package rest

import (
	"encoding/json"
	"github.com/asia-loop-gmbh/lambda-types-go/admin"
	"github.com/aws/aws-lambda-go/events"
)

func HTTPResponse(status int, request *events.APIGatewayProxyRequest, body interface{}) *events.APIGatewayProxyResponse {
	bodyString, err := json.Marshal(body)
	if err != nil {
		return HTTPErrorResponse(500, request, err)
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(bodyString),
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "*",
		},
	}
}

func HTTPErrorResponse(status int, request *events.APIGatewayProxyRequest, err error) *events.APIGatewayProxyResponse {
	body := admin.HttpErrorBody{
		Message:     err.Error(),
		Method:      request.HTTPMethod,
		Path:        request.Path,
		RequestID:   request.RequestContext.RequestID,
		RequestTime: request.RequestContext.RequestTime,
	}
	bodyString, err := json.Marshal(body)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: status,
			Body:       "{\"message\": \"Could not construct error object.\"}",
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "*",
			},
		}
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(bodyString),
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "*",
		},
	}
}
