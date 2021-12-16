package main

import (
	"encoding/json"
	"fmt"

	"github.com/scaleway/scaleway-functions-go/events"
	"github.com/scaleway/scaleway-functions-go/lambda"
	"klintt.io/detect/handlers/detectdaily"
)

type Model struct {
	Pairs   []string `json:"pairs"`
	OnlyFor string   `json:"onlyFor"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	model := Model{}
	err := json.Unmarshal([]byte(request.Body), &model)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{}, err
	}

	detectdaily.DetectOnManyPairToday(model.Pairs, model.OnlyFor, nil)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
