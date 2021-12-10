package main

import (
	"encoding/json"

	"github.com/scaleway/scaleway-functions-go/events"
	"github.com/scaleway/scaleway-functions-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	response := map[string]interface{}{
		"message": request.Body,
	}

	responseB, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseB),
		StatusCode: 201,
	}, nil
}

func main() {
	lambda.Start(handler)
}
