package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type body struct {
	Message string `json:"message"`
}

// Handler is used by AWS Lambda to handle request.
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Something")

	bStr, _ := json.Marshal(body{
		Message: "Everything OK!",
	})

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(bStr),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
