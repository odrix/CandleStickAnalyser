package main

import (
	"encoding/json"

	//"/handlers/detectdaily"

	"github.com/scaleway/scaleway-functions-go/events"
	"github.com/scaleway/scaleway-functions-go/lambda"
	"klintt.io/detect/handlers/detectdaily"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	response := map[string]interface{}{
		"message": request.Body,
	}

	responseB, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	pairs := []string{"BTCUSDT", "ETHUSDT", "BNBUSDT", "SOLUSDT", "ADAUSDT"}
	detectdaily.DetectOnManyPairToday(pairs, nil)

	return events.APIGatewayProxyResponse{
		Body:       string(responseB),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
